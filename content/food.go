package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Food content type
type Food struct {
	item.Header

	Title    string    `json:"title"`
	Cuisine  string    `json:"cuisine"`
	Category string    `json:"category"`
	Nonveg   bool      `json:"nonveg"`
	Body     string    `json:"body"`
	Size     string    `json:"size"`
	Serves   string    `json:"serves"`
	Price    string    `json:"price"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Food")] = func() interface{} { return new(Food) }
	item.Fields[strings.ToLower("Food")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Cuisine", Widget: item.WidgetInput, Helptext: "Cuisine type"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "Nonveg", Widget: item.WidgetInput, Helptext: "veg/non-veg"},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Size", Widget: item.WidgetInput, Helptext: "Enter the size here"},
		{Name: "Serves", Widget: item.WidgetInput, Helptext: "Serves persons"},
		{Name: "Price", Widget: item.WidgetInput, Helptext: "Enter the price here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Food is printed. Update it using more descriptive
// fields from the Food struct type
func (f *Food) String() string {
	return fmt.Sprintf("Food: %s", f.Title)
}

// Read Food
func (f *Food) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "food",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return f.Parse(resp.Content)
}

// FoodList1
func FoodList(lang, sortby string, size, skip int) ([]Food, int, error) {
	var foods []Food

	var req = &api.ListRequest{
		Type:     "food",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return foods, 0, err
	}

	for _, content := range resp.List {
		var f Food
		if err := f.Parse(content); err != nil {
			return foods, 0, err
		}
		foods = append(foods, f)
	}
	return foods, int(resp.Total), nil
}

// FoodSearch
func FoodSearch(lang, query string, size, skip int) ([]Food, int, error) {
	var foods []Food

	var req = &api.SearchRequest{
		Type:     "food",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return foods, 0, err
	}

	for _, content := range resp.Hits {
		var f Food
		if err := f.Parse(content); err != nil {
			return foods, 0, err
		}
		foods = append(foods, f)
	}
	return foods, int(resp.Total), nil
}

// Parse object
func (f *Food) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	f.Language = c["language"].(string)
	f.Slug = c["slug"].(string)

	f.ID = uint64(c["id"].(float64))
	f.Status = c["status"].(string)
	f.CreatedAt = int64(c["created_at"].(float64))
	f.UpdatedAt = int64(c["updated_at"].(float64))
	f.DeletedAt = int64(c["deleted_at"].(float64))

	f.Title = c["title"].(string)

	if _, ok := c["cuisine"]; ok {
		f.Cuisine = c["cuisine"].(string)
	}

	if _, ok := c["category"]; ok {
		f.Category = c["category"].(string)
	}

	if _, ok := c["nonveg"]; ok {
		f.Nonveg = c["nonveg"].(bool)
	}

	if _, ok := c["body"]; ok {
		f.Body = c["body"].(string)
	}

	if _, ok := c["size"]; ok {
		f.Size = c["size"].(string)
	}

	if _, ok := c["serves"]; ok {
		f.Serves = c["serves"].(string)
	}

	if _, ok := c["price"]; ok {
		f.Price = c["price"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			f.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			f.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			f.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			f.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			f.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			f.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
