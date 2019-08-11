package content

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Page content type
type Page struct {
	item.Header

	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Plaintext string    `json:"plaintext"`
	Richtext  string    `json:"richtext"`
	List      []string  `json:"list"`
	Image     item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Page")] = func() interface{} { return new(Page) }
	item.Fields[strings.ToLower("Page")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Subtitle", Widget: item.WidgetInput, Helptext: "Enter the Subtitle here"},
		{Name: "Plaintext", Widget: item.WidgetTextarea, Helptext: "Enter the plain text here"},
		{Name: "Richtext", Widget: item.WidgetRichtext, Helptext: "Enter the HTML content here"},
		{Name: "List", Widget: item.WidgetList, Helptext: "Add multple List items seperated by comma"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Page is printed. Update it using more descriptive
// fields from the Page struct type
func (p *Page) String() string {
	return fmt.Sprintf("Page: %s", p.Title)
}

// Read Page
func (p *Page) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "page",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	if resp.Err != "" {
		return errors.New(resp.Err)
	}

	return p.Parse(resp.Content)
}

// Get Page
func GetPage(lang, slug string) Page {
	var page Page
	page.Read(lang, slug)
	return page
}

// PageList1
func PageList(lang, sortby string, size, skip int) ([]Page, int, error) {
	var pages []Page

	var req = &api.ListRequest{
		Type:     "page",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return pages, 0, err
	}

	for _, content := range resp.List {
		var p Page
		if err := p.Parse(content); err != nil {
			return pages, 0, err
		}
		pages = append(pages, p)
	}
	return pages, int(resp.Total), nil
}

// PageSearch
func PageSearch(lang, query string, size, skip int) ([]Page, int, error) {
	var pages []Page

	var req = &api.SearchRequest{
		Type:     "page",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return pages, 0, err
	}

	for _, content := range resp.Hits {
		var p Page
		if err := p.Parse(content); err != nil {
			return pages, 0, err
		}
		pages = append(pages, p)
	}
	return pages, int(resp.Total), nil
}

// Parse object
func (p *Page) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	p.Language = c["language"].(string)
	p.Slug = c["slug"].(string)

	p.ID = uint64(c["id"].(float64))
	p.Status = c["status"].(string)
	p.CreatedAt = int64(c["created_at"].(float64))
	p.UpdatedAt = int64(c["updated_at"].(float64))
	p.DeletedAt = int64(c["deleted_at"].(float64))

	p.Title = c["title"].(string)

	if _, ok := c["subtitle"]; ok {
		p.Subtitle = c["subtitle"].(string)
	}
	if _, ok := c["plaintext"]; ok {
		p.Plaintext = c["plaintext"].(string)
	}
	if _, ok := c["richtext"]; ok {
		p.Richtext = c["richtext"].(string)
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
