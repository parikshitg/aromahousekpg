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

	Title            string    `json:"title"`
	NumNights        string    `json:"num_nights"`
	NumAdults        string    `json:"num_adults"`
	NumChilds        string    `json:"num_childs"`
	Charges          string    `json:"charges"`
	ExtraChildCharge string    `json:"extra_child_charge"`
	ExtraAdultCharge string    `json:"extra_adult_charge"`
	MealPlan         string    `json:"meal_plan"`
	Body             string    `json:"body"`
	Image            item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Package")] = func() interface{} { return new(Package) }
	item.Fields[strings.ToLower("Package")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "NumNights", Heading: " ", HasLabel: true, Widget: item.WidgetInput, Helptext: "num of nights", SkipFooter: true},
		{Name: "NumAdults", HasLabel: true, Widget: item.WidgetInput, Helptext: "No. of adults", SkipHeader: true, SkipFooter: true},
		{Name: "NumChilds", HasLabel: true, Widget: item.WidgetInput, Helptext: "No. of Childs", SkipHeader: true},
		{Name: "Charges", Heading: " ", HasLabel: true, Widget: item.WidgetInput, Helptext: "charges", SkipFooter: true},
		{Name: "ExtraChildCharge", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra Child Charge", SkipHeader: true, SkipFooter: true},
		{Name: "ExtraAdultCharge", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra Adult Charge", SkipHeader: true},
		{Name: "MealPlan", Widget: item.WidgetInput, Helptext: "CP, AP, MAP or EP"},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
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

	if _, ok := c["num_nights"]; ok {
		p.NumNights = c["num_nights"].(string)
	}

	if _, ok := c["num_adults"]; ok {
		p.NumAdults = c["num_adults"].(string)
	}

	if _, ok := c["num_childs"]; ok {
		p.NumChilds = c["num_childs"].(string)
	}

	if _, ok := c["charges"]; ok {
		p.Charges = c["charges"].(string)
	}

	if _, ok := c["extra_adult_charge"]; ok {
		p.ExtraAdultCharge = c["extra_adult_charge"].(string)
	}

	if _, ok := c["extra_child_charge"]; ok {
		p.ExtraChildCharge = c["extra_child_charge"].(string)
	}

	if _, ok := c["meal_plan"]; ok {
		p.MealPlan = c["meal_plan"].(string)
	}

	if _, ok := c["body"]; ok {
		p.Body = c["body"].(string)
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
