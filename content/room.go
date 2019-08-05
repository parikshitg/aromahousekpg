package content

import (
	"context"
	"fmt"
	"strings"

	"git.urantiatech.com/cloudcms/cloudcms/api"
	"git.urantiatech.com/cloudcms/cloudcms/item"
)

// Room content type
type Room struct {
	item.Header

	Title    string    `json:"title"`
	BedType  string    `json:"bed_type"`
	NumBeds  string    `json:"num_beds"`
	Capacity string    `json:"capacity"`
	Body     string    `json:"body"`
	Charges  string    `json:"charges"`
	Category string    `json:"category"`
	Image    item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Room")] = func() interface{} { return new(Room) }
	item.Fields[strings.ToLower("Room")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},
		{Name: "BedType", Widget: item.WidgetInput, Helptext: "Master Bed, Queen Bed, Single bed, Matress etc"},
		{Name: "NumBeds", Widget: item.WidgetInput, Helptext: "Number of Beds"},
		{Name: "Capacity", Widget: item.WidgetInput, Helptext: "Max number of Adults"},
		{Name: "Body", Widget: item.WidgetRichtext, Helptext: "Enter the Body here"},
		{Name: "Charges", Widget: item.WidgetInput, Helptext: "Room charges"},
		{Name: "Category", Widget: item.WidgetInput, Helptext: "Enter the Category here"},
		{Name: "file:Image", Widget: item.WidgetFile, Helptext: "Select Image", FileType: item.FileImageType},
	}...)
}

// String defines how a Room is printed. Update it using more descriptive
// fields from the Room struct type
func (r *Room) String() string {
	return fmt.Sprintf("Room: %s", r.Title)
}

// Read Room
func (r *Room) Read(lang, slug string) error {
	var req = &api.ReadRequest{
		Type:     "room",
		Language: lang,
		Slug:     slug,
	}

	resp, err := svc.Read(context.Background(), req)
	if err != nil {
		return err
	}

	return r.Parse(resp.Content)
}

// RoomList1
func RoomList(lang, sortby string, size, skip int) ([]Room, int, error) {
	var rooms []Room

	var req = &api.ListRequest{
		Type:     "room",
		Language: lang,
		SortBy:   sortby,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.List(context.Background(), req)
	if err != nil {
		return rooms, 0, err
	}

	for _, content := range resp.List {
		var r Room
		if err := r.Parse(content); err != nil {
			return rooms, 0, err
		}
		rooms = append(rooms, r)
	}
	return rooms, int(resp.Total), nil
}

// RoomSearch
func RoomSearch(lang, query string, size, skip int) ([]Room, int, error) {
	var rooms []Room

	var req = &api.SearchRequest{
		Type:     "room",
		Language: lang,
		Query:    query,
		Size:     size,
		Skip:     skip,
	}

	resp, err := svc.Search(context.Background(), req)
	if err != nil {
		return rooms, 0, err
	}

	for _, content := range resp.Hits {
		var r Room
		if err := r.Parse(content); err != nil {
			return rooms, 0, err
		}
		rooms = append(rooms, r)
	}
	return rooms, int(resp.Total), nil
}

// Parse object
func (r *Room) Parse(contents interface{}) error {
	c := contents.(map[string]interface{})

	r.Language = c["language"].(string)
	r.Slug = c["slug"].(string)

	r.ID = uint64(c["id"].(float64))
	r.Status = c["status"].(string)
	r.CreatedAt = int64(c["created_at"].(float64))
	r.UpdatedAt = int64(c["updated_at"].(float64))
	r.DeletedAt = int64(c["deleted_at"].(float64))

	r.Title = c["title"].(string)

	if _, ok := c["bed_type"]; ok {
		r.BedType = c["bed_type"].(string)
	}

	if _, ok := c["num_beds"]; ok {
		r.NumBeds = c["num_beds"].(string)
	}

	if _, ok := c["capacity"]; ok {
		r.Capacity = c["capacity"].(string)
	}

	if _, ok := c["body"]; ok {
		r.Body = c["body"].(string)
	}

	if _, ok := c["charges"]; ok {
		r.Charges = c["charges"].(string)
	}

	if _, ok := c["category"]; ok {
		r.Category = c["category"].(string)
	}

	if v, ok := c["file:image"]; ok {
		image := v.(map[string]interface{})
		if _, ok := image["name"]; ok {
			r.Image.Name = image["name"].(string)
		}
		if _, ok := image["size"]; ok {
			r.Image.Size = int64(image["size"].(float64))
		}
		if _, ok := image["uri"]; ok {
			r.Image.URI = image["uri"].(string)
		}
	} else {
		if _, ok := c["file:image.name"]; ok {
			r.Image.Name = c["file:image.name"].(string)
		}
		if _, ok := c["file:image.size"]; ok {
			r.Image.Size = int64(c["file:image.size"].(float64))
		}
		if _, ok := c["file:image.uri"]; ok {
			r.Image.URI = c["file:image.uri"].(string)
		}
	}

	return nil
}
