package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Testimonial content type
type Testimonial struct {
	item.Header

	Guest string    `json:"guest"`
	Text  string    `json:"text"`
	Image item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Testimonial")] = func() interface{} { return new(Testimonial) }
	item.Fields[strings.ToLower("Testimonial")] = append(item.HeaderFields, []item.Field{
		{Name: "Guest", Widget: item.WidgetInput, Helptext: "Enter the Guest Name here", UseForSlug: true},
		{Name: "Text", Widget: item.WidgetRichtext, Helptext: "Enter the Text here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Testimonial is printed. Update it using more descriptive
// fields from the Testimonial struct type
func (t *Testimonial) String() string {
	return fmt.Sprintf("Testimonial: %s", t.Guest)
}

// Read Testimonial
func (t *Testimonial) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "review",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return t.Parse(resp.Content)
}

// TestimonialList1
func TestimonialList(lang, sortby string, size, skip int) ([]Testimonial, int, error) {
	var reviews []Testimonial

	var req = &api.ListRequest{
		Type:     "review",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return reviews, 0, err
	}

	for _, content := range resp.List {
		var t Testimonial
		if err := t.Parse(content); err != nil {
			return reviews, 0, err
		}
		reviews = append(reviews, t)
	}
	return reviews, int(resp.Total), nil
}

// TestimonialSearch
func TestimonialSearch(lang, query string, size, skip int) ([]Testimonial, int, error) {
	var reviews []Testimonial

	var req = &api.SearchRequest{
		Type:     "review",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return reviews, 0, err
	}

	for _, content := range resp.Hits {
		var t Testimonial
		if err := t.Parse(content); err != nil {
			return reviews, 0, err
		}
		reviews = append(reviews, t)
	}
	return reviews, int(resp.Total), nil
}

// Parse object
func (t *Testimonial) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	t.Language = c["language"].(string)
	t.Slug = c["slug"].(string)

	t.ID = uint64(c["id"].(float64))
	t.Status = c["status"].(string)
	t.CreatedAt = int64(c["created_at"].(float64))
	t.UpdatedAt = int64(c["updated_at"].(float64))
	t.DeletedAt = int64(c["deleted_at"].(float64))

	t.Guest = c["guest"].(string)

	if _, ok := c["text"]; ok {
		t.Text = c["text"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			t.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			t.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			t.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			t.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			t.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			t.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
