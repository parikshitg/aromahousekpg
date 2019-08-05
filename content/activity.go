package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Activity content type
type Activity struct {
	item.Header

	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Image    item.File `json:"file:image"`
}

//var svc = s.Service{}

func init() {
	item.Types[strings.ToLower("Activity")] = func() interface{} { return new(Activity) }
	item.Fields[strings.ToLower("Activity")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "Tags", Widget: item.WidgetTags, Helptext: "Add multple Tags seperated by comma"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Activity is printed. Update it using more descriptive
// fields from the Activity struct type
func (a *Activity) String() string {
	return fmt.Sprintf("Activity: %s", a.Title)
}

// Read Activity
func (a *Activity) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "activity",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return a.Parse(resp.Content)
}

// ActivityList1
func ActivityList(lang, sortby string, size, skip int) ([]Activity, int, error) {
	var activities []Activity

	var req = &api.ListRequest{
		Type:     "activity",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return activities, 0, err
	}

	for _, content := range resp.List {
		var a Activity
		if err := a.Parse(content); err != nil {
			return activities, 0, err
		}
		activities = append(activities, a)
	}
	return activities, int(resp.Total), nil
}

// ActivitySearch
func ActivitySearch(lang, query string, size, skip int) ([]Activity, int, error) {
	var activities []Activity

	var req = &api.SearchRequest{
		Type:     "activity",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return activities, 0, err
	}

	for _, content := range resp.Hits {
		var a Activity
		if err := a.Parse(content); err != nil {
			return activities, 0, err
		}
		activities = append(activities, a)
	}
	return activities, int(resp.Total), nil
}

// Parse object
func (a *Activity) Parse(contents interface{}) error {
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
