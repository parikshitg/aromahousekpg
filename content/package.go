package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Package content type
type Package struct {
	item.Header

	Title     string    `json:"title"`
	RoomType  string    `json:"roomtype"`
	NumNights string    `json:"numnights"`
	NumAdults string    `json:"numadults"`
	NumChilds string    `json:"numchilds"`
	Charges   string    `json:"charges"`
	Details   string    `json:"details"`
	List      []string  `json:"list"`
	Image     item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Package")] = func() interface{} { return new(Package) }
	item.Fields[strings.ToLower("Package")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "RoomType", Widget: item.WidgetInput, Helptext: "Room type", SkipFooter: true},
		{Name: "NumNights", HasLabel: true, Widget: item.WidgetInput, Helptext: "Num of nights", SkipHeader: true},
		{Name: "NumAdults", Widget: item.WidgetInput, Helptext: "No. of adults", SkipFooter: true},
		{Name: "NumChilds", HasLabel: true, Widget: item.WidgetInput, Helptext: "No. of Childs", SkipHeader: true},
		{Name: "Charges", Widget: item.WidgetInput, Helptext: "charges"},
		{Name: "Details", Widget: item.WidgetRichtext, Helptext: "Enter the Details here"},
		{Name: "List", Widget: item.WidgetList, Helptext: "List"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Package is printed. Update it using more descriptive
// fields from the Package struct type
func (p *Package) String() string {
	return fmt.Sprintf("Package: %s", p.Title)
}

// Read Package
func (p *Package) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "package",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return p.Parse(resp.Content)
}

// PackageList1
func PackageList(lang, sortby string, size, skip int) ([]Package, int, error) {
	var packages []Package

	var req = &api.ListRequest{
		Type:     "package",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return packages, 0, err
	}

	for _, content := range resp.List {
		var p Package
		if err := p.Parse(content); err != nil {
			return packages, 0, err
		}
		packages = append(packages, p)
	}
	return packages, int(resp.Total), nil
}

// PackageSearch
func PackageSearch(lang, query string, size, skip int) ([]Package, int, error) {
	var packages []Package

	var req = &api.SearchRequest{
		Type:     "package",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return packages, 0, err
	}

	for _, content := range resp.Hits {
		var p Package
		if err := p.Parse(content); err != nil {
			return packages, 0, err
		}
		packages = append(packages, p)
	}
	return packages, int(resp.Total), nil
}

// Parse object
func (p *Package) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	p.Language = c["language"].(string)
	p.Slug = c["slug"].(string)

	p.ID = uint64(c["id"].(float64))
	p.Status = c["status"].(string)
	p.CreatedAt = int64(c["created_at"].(float64))
	p.UpdatedAt = int64(c["updated_at"].(float64))
	p.DeletedAt = int64(c["deleted_at"].(float64))

	p.Title = c["title"].(string)

	if _, ok := c["roomtype"]; ok {
		p.RoomType = c["roomtype"].(string)
	}

	if _, ok := c["numnights"]; ok {
		p.NumNights = c["numnights"].(string)
	}

	if _, ok := c["numadults"]; ok {
		p.NumAdults = c["numadults"].(string)
	}

	if _, ok := c["numchilds"]; ok {
		p.NumChilds = c["numchilds"].(string)
	}

	if _, ok := c["charges"]; ok {
		p.Charges = c["charges"].(string)
	}

	if _, ok := c["details"]; ok {
		p.Details = c["details"].(string)
	}

	if v, ok := c["list"]; ok && v != nil {
		switch c["list"].(type) {
		case string:
			p.List = []string{c["list"].(string)}
		case []interface{}:
			vals := c["list"].([]interface{})
			for _, val := range vals {
				p.List = append(p.List, val.(string))
			}
		}
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			p.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			p.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			p.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			p.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			p.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			p.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
