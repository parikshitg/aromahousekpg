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

	Title         string    `json:"title"`
	RoomType      string    `json:roomtype"`
	Capacity      string    `json:"capacity"`
	DoubleBeds    string    `json:"doublebeds"`
	SingleBeds    string    `json:"singlebeds"`
	Matress       string    `json:"matress"`
	Details       string    `json:"details"`
	CPCharges     string    `json:"cpcharges"`
	CPExtraAdult  string    `json:"cpextraadult"`
	CPExtraChild  string    `json:"cpextrachild"`
	EPCharges     string    `json:"epcharges"`
	EPExtraAdult  string    `json:"epextraadult"`
	EPExtraChild  string    `json:"epextrachild"`
	APCharges     string    `json:"apcharges"`
	APExtraAdult  string    `json:"apextraadult"`
	APExtraChild  string    `json:"apextrachild"`
	MAPCharges    string    `json:"mapcharges"`
	MAPExtraAdult string    `json:"mapextraadult"`
	MAPExtraChild string    `json:"mapextrachild"`
	Facilities    []string  `json:"facilities"`
	GST           string    `json:"gst"`
	Weight        string    `json:"weight"`
	Image         item.File `json:"file:image"`
}

func init() {
	item.Types[strings.ToLower("Room")] = func() interface{} { return new(Room) }
	item.Fields[strings.ToLower("Room")] = append(item.HeaderFields, []item.Field{
		{Name: "Title", Widget: item.WidgetInput, Helptext: "Enter the Title here", UseForSlug: true},

		{Name: "RoomType", Widget: item.WidgetInput, Helptext: "Type of room", SkipFooter: true},
		{Name: "Capacity", Widget: item.WidgetInput, HasLabel: true, Helptext: "Max number of persons", SkipHeader: true},

		{Name: "DoubleBeds", Heading: "Num Beds", HasLabel: true, Widget: item.WidgetInput, Helptext: "Number of Double Beds", SkipFooter: true},
		{Name: "SingleBeds", HasLabel: true, Widget: item.WidgetInput, Helptext: "Number of Single Beds", SkipHeader: true, SkipFooter: true},
		{Name: "Matress", HasLabel: true, Widget: item.WidgetInput, Helptext: "Number of Matresses", SkipHeader: true},

		{Name: "Details", Widget: item.WidgetRichtext, Helptext: "Enter the room details here"},

		{Name: "CPCharges", Heading: "CP", HasLabel: true, Widget: item.WidgetInput, Helptext: "Room charges", SkipFooter: true},
		{Name: "CPExtraAdult", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra adult charges", SkipHeader: true, SkipFooter: true},
		{Name: "CPExtraChild", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra child charges", SkipHeader: true},

		{Name: "EPCharges", Heading: "EP", HasLabel: true, Widget: item.WidgetInput, Helptext: "Room charges", SkipFooter: true},
		{Name: "EPExtraAdult", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra adult charges", SkipHeader: true, SkipFooter: true},
		{Name: "EPExtraChild", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra child charges", SkipHeader: true},

		{Name: "APCharges", Heading: "AP", HasLabel: true, Widget: item.WidgetInput, Helptext: "Room charges", SkipFooter: true},
		{Name: "APExtraAdult", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra adult charges", SkipHeader: true, SkipFooter: true},
		{Name: "APExtraChild", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra child charges", SkipHeader: true},

		{Name: "MAPCharges", Heading: "MAP", HasLabel: true, Widget: item.WidgetInput, Helptext: "Room charges", SkipFooter: true},
		{Name: "MAPExtraAdult", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra adult charges", SkipHeader: true, SkipFooter: true},
		{Name: "MAPExtraChild", HasLabel: true, Widget: item.WidgetInput, Helptext: "Extra child charges", SkipHeader: true},

		{Name: "Facilities", Widget: item.WidgetList, Helptext: "Add multple Facility seperated by comma"},
		{Name: "GST", Widget: item.WidgetInput, Helptext: "Enter the GST percentage here"},
		{Name: "Weight", Widget: item.WidgetInput, Helptext: "Enter the Weight here"},
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

// Get Room
func GetRoom(lang, slug string) Room {
	var room Room
	room.Read(lang, slug)
	return room
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

	if _, ok := c["roomtype"]; ok {
		r.RoomType = c["roomtype"].(string)
	}

	if _, ok := c["capacity"]; ok {
		r.Capacity = c["capacity"].(string)
	}

	if _, ok := c["doublebeds"]; ok {
		r.DoubleBeds = c["doublebeds"].(string)
	}

	if _, ok := c["singlebeds"]; ok {
		r.SingleBeds = c["singlebeds"].(string)
	}

	if _, ok := c["matress"]; ok {
		r.Matress = c["matress"].(string)
	}

	if _, ok := c["details"]; ok {
		r.Details = c["details"].(string)
	}

	// CP
	if _, ok := c["cpcharges"]; ok {
		r.CPCharges = c["cpcharges"].(string)
	}

	if _, ok := c["cpextraadult"]; ok {
		r.CPExtraAdult = c["cpextraadult"].(string)
	}

	if _, ok := c["cpextrachild"]; ok {
		r.CPExtraChild = c["cpextrachild"].(string)
	}

	// EP
	if _, ok := c["epcharges"]; ok {
		r.EPCharges = c["epcharges"].(string)
	}

	if _, ok := c["epextraadult"]; ok {
		r.EPExtraAdult = c["epextraadult"].(string)
	}

	if _, ok := c["epextrachild"]; ok {
		r.EPExtraChild = c["epextrachild"].(string)
	}

	// AP
	if _, ok := c["apcharges"]; ok {
		r.APCharges = c["apcharges"].(string)
	}

	if _, ok := c["apextraadult"]; ok {
		r.APExtraAdult = c["apextraadult"].(string)
	}

	if _, ok := c["apextrachild"]; ok {
		r.APExtraChild = c["apextrachild"].(string)
	}

	// MAP
	if _, ok := c["cpcharges"]; ok {
		r.MAPCharges = c["mapcharges"].(string)
	}

	if _, ok := c["mapextraadult"]; ok {
		r.MAPExtraAdult = c["mapextraadult"].(string)
	}

	if _, ok := c["mapextrachild"]; ok {
		r.MAPExtraChild = c["mapextrachild"].(string)
	}

	if v, ok := c["facilities"]; ok && v != nil {
		switch c["facilities"].(type) {
		case string:
			r.Facilities = []string{c["facilities"].(string)}
		case []interface{}:
			vals := c["facilities"].([]interface{})
			for _, val := range vals {
				r.Facilities = append(r.Facilities, val.(string))
			}
		}
	}

	if _, ok := c["gst"]; ok {
		r.GST = c["gst"].(string)
	}

	if _, ok := c["weight"]; ok {
		r.Weight = c["weight"].(string)
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
