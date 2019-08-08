package content

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Banner content type
type Banner struct {
	item.Header

	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Text     string    `json:"text"`
	Weight   string    `json:"weight"`
	Link     string    `json:"link"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Banner")] = func() interface{} { return new(Banner) }
	item.Fields[strings.ToLower("Banner")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Subtitle", Widget: item.WidgetInput, Helptext: "Enter the Subtitle here"},
		{Name: "Text", Widget: item.WidgetInput, Helptext: "Enter the Text here"},
		{Name: "Weight", Widget: item.WidgetInput, Helptext: "Enter the Weight here"},
		{Name: "Link", Widget: item.WidgetInput, Helptext: "Enter the page URI here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Banner is printed. Update it using more descriptive
// fields from the Banner struct type
func (b *Banner) String() string {
	return fmt.Sprintf("Banner: %s", b.Title)
}

// Content converts Banner to content
func (b *Banner) Content() map[string]interface{} {
	var content map[string]interface{}
	bannerString, _ := json.Marshal(b)
	json.Unmarshal(bannerString, &content)
	return content
}

// Create Banner
func (b *Banner) Create(lang, slugtext string) error {
	var req = &api.CreateRequest{
		Type:     "banner",
		Language: lang,
		SlugText: slugtext,
		Content:  b.Content(),
	}

	_, err := svc.Create(context.Background(), req, false)
	return err
}

// Read Banner
func (b *Banner) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "banner",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return b.Parse(resp.Content)
}

// Update Banner
func (b *Banner) Update(lang, slug string) error {
	var req = &api.UpdateRequest{
		Type:     "banner",
		Language: lang,
		Slug:     slug,
		Content:  b.Content(),
	}

	_, err := svc.Update(context.Background(), req, false)
	return err
}

// Delete Banner
func (b *Banner) Delete(lang, slug string) error {
	var req = &api.DeleteRequest{
		Type:     "banner",
		Language: lang,
		Slug:     slug,
	}

	_, err := svc.Delete(context.Background(), req, false)
	if err != nil {
		return err
	}

	return nil
}

// BannerList1
func BannerList(lang, sortby string, size, skip int) ([]Banner, int, error) {
	var banners []Banner

	var req = &api.ListRequest{
		Type:     "banner",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return banners, 0, err
	}

	for _, content := range resp.List {
		var b Banner
		if err := b.Parse(content); err != nil {
			return banners, 0, err
		}
		banners = append(banners, b)
	}
	return banners, int(resp.Total), nil
}

// Parse object
func (b *Banner) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	b.Language = c["language"].(string)
	b.Slug = c["slug"].(string)

	b.ID = uint64(c["id"].(float64))
	b.Status = c["status"].(string)
	b.CreatedAt = int64(c["created_at"].(float64))
	b.UpdatedAt = int64(c["updated_at"].(float64))
	b.DeletedAt = int64(c["deleted_at"].(float64))

	b.Title = c["title"].(string)

	if _, ok := c["subtitle"]; ok {
		b.Subtitle = c["subtitle"].(string)
	}

	if _, ok := c["text"]; ok {
		b.Text = c["text"].(string)
	}

	if _, ok := c["weight"]; ok {
		b.Weight = c["weight"].(string)
	}

	if _, ok := c["link"]; ok {
		b.Link = c["link"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			b.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			b.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			b.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			b.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			b.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			b.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
