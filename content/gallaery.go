package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Gallery content type
type Gallery struct {
	item.Header

	Title string    `json:"title"`
	Text  string    `json:"text"`
	Image item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Gallery")] = func() interface{} { return new(Gallery) }
	item.Fields[strings.ToLower("Gallery")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Heading: "Guest", Widget: item.WidgetInput, Helptext: "Image Title", UseForSlug: true},
		{Name: "Text", Widget: item.WidgetTextarea, Helptext: "Enter the Text here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Gallery is printed. Update it using more descriptive
// fields from the Gallery struct type
func (g *Gallery) String() string {
	return fmt.Sprintf("Gallery: %s", g.Title)
}

// Read Gallery
func (g *Gallery) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "gallery",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return g.Parse(resp.Content)
}

// GalleryList1
func GalleryList(lang, sortby string, size, skip int) ([]Gallery, int, error) {
	var galleries []Gallery

	var req = &api.ListRequest{
		Type:     "gallery",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return galleries, 0, err
	}

	for _, content := range resp.List {
		var g Gallery
		if err := g.Parse(content); err != nil {
			return galleries, 0, err
		}
		galleries = append(galleries, g)
	}
	return galleries, int(resp.Total), nil
}

// GallerySearch
func GallerySearch(lang, query string, size, skip int) ([]Gallery, int, error) {
	var galleries []Gallery

	var req = &api.SearchRequest{
		Type:     "gallery",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return galleries, 0, err
	}

	for _, content := range resp.Hits {
		var g Gallery
		if err := g.Parse(content); err != nil {
			return galleries, 0, err
		}
		galleries = append(galleries, g)
	}
	return galleries, int(resp.Total), nil
}

// Parse object
func (g *Gallery) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	g.Language = c["language"].(string)
	g.Slug = c["slug"].(string)

	g.ID = uint64(c["id"].(float64))
	g.Status = c["status"].(string)
	g.CreatedAt = int64(c["created_at"].(float64))
	g.UpdatedAt = int64(c["updated_at"].(float64))
	g.DeletedAt = int64(c["deleted_at"].(float64))

	if _, ok := c["title"]; ok {
		g.Title = c["title"].(string)
	}

	if _, ok := c["text"]; ok {
		g.Text = c["text"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			g.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			g.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			g.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			g.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			g.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			g.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
