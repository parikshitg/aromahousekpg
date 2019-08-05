package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Drink content type
type Drink struct {
	item.Header

	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Category string    `json:"category"`
	Price    string    `json:"price"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Drink")] = func() interface{} { return new(Drink) }
	item.Fields[strings.ToLower("Drink")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "Price", Widget: item.WidgetInput, Helptext: "Enter the price here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Drink is printed. Update it using more descriptive
// fields from the Drink struct type
func (d *Drink) String() string {
	return fmt.Sprintf("Drink: %s", d.Title)
}

// Read Drink
func (d *Drink) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "drink",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return d.Parse(resp.Content)
}

// DrinkList1
func DrinkList(lang, sortby string, size, skip int) ([]Drink, int, error) {
	var drinks []Drink

	var req = &api.ListRequest{
		Type:     "drink",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return drinks, 0, err
	}

	for _, content := range resp.List {
		var d Drink
		if err := d.Parse(content); err != nil {
			return drinks, 0, err
		}
		drinks = append(drinks, d)
	}
	return drinks, int(resp.Total), nil
}

// DrinkSearch
func DrinkSearch(lang, query string, size, skip int) ([]Drink, int, error) {
	var drinks []Drink

	var req = &api.SearchRequest{
		Type:     "drink",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return drinks, 0, err
	}

	for _, content := range resp.Hits {
		var d Drink
		if err := d.Parse(content); err != nil {
			return drinks, 0, err
		}
		drinks = append(drinks, d)
	}
	return drinks, int(resp.Total), nil
}

// Parse object
func (d *Drink) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	d.Language = c["language"].(string)
	d.Slug = c["slug"].(string)

	d.ID = uint64(c["id"].(float64))
	d.Status = c["status"].(string)
	d.CreatedAt = int64(c["created_at"].(float64))
	d.UpdatedAt = int64(c["updated_at"].(float64))
	d.DeletedAt = int64(c["deleted_at"].(float64))

	d.Title = c["title"].(string)

	if _, ok := c["body"]; ok {
		d.Body = c["body"].(string)
	}

	if _, ok := c["category"]; ok {
		d.Category = c["category"].(string)
	}

	if _, ok := c["price"]; ok {
		d.Price = c["price"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			d.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			d.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			d.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			d.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			d.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			d.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
