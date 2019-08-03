package content

import (
	"fmt"
	"strings"
	"time"

	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Blog content type
type Blog struct {
	item.Header

	Title    string    `json:"title"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	Body     string    `json:"body"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Blog")] = func() interface{} { return new(Blog) }
	item.Fields[strings.ToLower("Blog")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "Author", Heading: " ", HasLabel: true, Widget: item.WidgetInput, Helptext: "Author Name", SkipFooter: true},
		{Name: "Date", HasLabel: true, Widget: item.WidgetDate, Helptext: "YYYY-MM-DD", SkipHeader: true},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "Tags", Widget: item.WidgetTags, Helptext: "Add multple Tags seperated by comma"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Blog is printed. Update it using more descriptive
// fields from the Blog struct type
func (b *Blog) String() string {
	return fmt.Sprintf("This is Blog %s", b.Title)
}

// Parse object
func (b *Blog) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	b.Language = c["language"].(string)
	b.Slug = c["slug"].(string)

	b.ID = uint64(c["id"].(float64))
	b.Status = c["status"].(string)
	b.CreatedAt = int64(c["created_at"].(float64))
	b.UpdatedAt = int64(c["updated_at"].(float64))
	b.DeletedAt = int64(c["deleted_at"].(float64))

	b.Title = c["title"].(string)

	if _, ok := c["author"]; ok {
		b.Author = c["author"].(string)
	}

	if _, ok := c["date"]; ok {
		b.Date, _ = time.Parse(time.RFC3339, c["date"].(string))
	}

	if _, ok := c["body"]; ok {
		b.Body = c["body"].(string)
	}

	if _, ok := c["category"]; ok {
		b.Category = c["category"].(string)
	}

	if v, ok := c["tags"]; ok && v != nil {
		switch c["tags"].(type) {
		case string:
			b.Tags = []string{c["tags"].(string)}
		case []interface{}:
			vals := c["tags"].([]interface{})
			for _, val := range vals {
				b.Tags = append(b.Tags, val.(string))
			}
		}
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
