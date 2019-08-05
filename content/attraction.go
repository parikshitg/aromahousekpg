package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Attraction content type
type Attraction struct {
	item.Header

	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Location string    `json:"location"`
	Distance string    `json:"distance"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Attraction")] = func() interface{} { return new(Attraction) }
	item.Fields[strings.ToLower("Attraction")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Location", Widget: item.WidgetInput, Helptext: "Enter the Location here"},
		{Name: "Distance", Widget: item.WidgetInput, Helptext: "Enter the Distance here"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "Tags", Widget: item.WidgetTags, Helptext: "Add multple Tags seperated by comma"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Attraction is printed. Update it using more descriptive
// fields from the Attraction struct type
func (a *Attraction) String() string {
	return fmt.Sprintf("Attraction: %s", a.Title)
}

// Read Attraction
func (a *Attraction) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "attraction",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return a.Parse(resp.Content)
}

// AttractionList1
func AttractionList(lang, sortby string, size, skip int) ([]Attraction, int, error) {
	var attractions []Attraction

	var req = &api.ListRequest{
		Type:     "attraction",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return attractions, 0, err
	}

	for _, content := range resp.List {
		var a Attraction
		if err := a.Parse(content); err != nil {
			return attractions, 0, err
		}
		attractions = append(attractions, a)
	}
	return attractions, int(resp.Total), nil
}

// AttractionSearch
func AttractionSearch(lang, query string, size, skip int) ([]Attraction, int, error) {
	var attractions []Attraction

	var req = &api.SearchRequest{
		Type:     "attraction",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return attractions, 0, err
	}

	for _, content := range resp.Hits {
		var a Attraction
		if err := a.Parse(content); err != nil {
			return attractions, 0, err
		}
		attractions = append(attractions, a)
	}
	return attractions, int(resp.Total), nil
}

// Parse object
func (a *Attraction) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	a.Language = c["language"].(string)
	a.Slug = c["slug"].(string)

	a.ID = uint64(c["id"].(float64))
	a.Status = c["status"].(string)
	a.CreatedAt = int64(c["created_at"].(float64))
	a.UpdatedAt = int64(c["updated_at"].(float64))
	a.DeletedAt = int64(c["deleted_at"].(float64))

	a.Title = c["title"].(string)

	if _, ok := c["body"]; ok {
		a.Body = c["body"].(string)
	}

	if _, ok := c["location"]; ok {
		a.Location = c["location"].(string)
	}

	if _, ok := c["distance"]; ok {
		a.Distance = c["distance"].(string)
	}

	if _, ok := c["category"]; ok {
		a.Category = c["category"].(string)
	}

	if v, ok := c["tags"]; ok && v != nil {
		switch c["tags"].(type) {
		case string:
			a.Tags = []string{c["tags"].(string)}
		case []interface{}:
			vals := c["tags"].([]interface{})
			for _, val := range vals {
				a.Tags = append(a.Tags, val.(string))
			}
		}
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			a.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			a.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			a.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			a.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			a.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			a.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
