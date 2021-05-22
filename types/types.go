package types

type RSS3ID = string
type RSS3ItemID = string
type RSS3ItemsID = string
type RSS3ListID = string
type ThirdPartyAddress = []string

// RSS3Base Common attributes for each file
type RSS3Base struct {
	//Id          interface{} `json:"id"` // RSS3ID | RSS3ItemsID | RSS3ListID
	Version     string `json:"@version"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Signature   string `json:"signature,omitempty"`
}

// RSS3 file, Entrance
type RSS3 struct {
	*RSS3Base

	Id RSS3ID `json:"id"`

	Profile struct {
		Name      string            `json:"name,omitempty"`
		Avatar    ThirdPartyAddress `json:"avatar,omitempty"`
		Bio       string            `json:"bio,omitempty"`
		Tags      []string          `json:"tags,omitempty"`
		Signature string            `json:"signature,omitempty"`
	} `json:"profile"`

	Items     []RSS3Item  `json:"items"`
	ItemsNext RSS3ItemsID `json:"items_next,omitempty"`

	Links []struct {
		Type      string     `json:"type"`
		Tags      []string   `json:"tags,omitempty"`
		List      []RSS3ID   `json:"list"`
		ListNext  RSS3ListID `json:"list_next,omitempty"`
		Signature string     `json:"signature,omitempty"`
	} `json:"links,omitempty"`

	Backlinks []struct {
		Type     string     `json:"type"`
		List     []RSS3ID   `json:"list"`
		ListNext RSS3ListID `json:"list_next,omitempty"`
	} `json:"@backlinks,omitempty"`

	Assets []struct {
		Type    string   `json:"type"`
		Tags    []string `json:"tags,omitempty"`
		Content string   `json:"content"`
	} `json:"assets,omitempty"`
}

// RSS3Items file
type RSS3Items struct {
	*RSS3Base

	Id RSS3ItemsID `json:"id"`

	Items     []RSS3Item  `json:"items"`
	ItemsNext RSS3ItemsID `json:"items_next,omitempty"`
}

// RSS3List file
type RSS3List struct {
	*RSS3Base

	Id RSS3ListID `json:"id"`

	List []interface{} `json:"list"` // [] RSS3ID | [] RSS3ItemID

	ListNext RSS3ListID `json:"list_next,omitempty"`
}

// RSS3Item file
type RSS3Item struct {
	Id            string   `json:"id"`
	Authors       []RSS3ID `json:"authors,omitempty"`
	Title         string   `json:"title,omitempty"`
	Summary       string   `json:"summary,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	DatePublished string   `json:"date_published,omitempty"`
	DateModified  string   `json:"date_modified,omitempty"`

	Type     string     `json:"type,omitempty"`
	Upstream RSS3ItemID `json:"upstream,omitempty"`

	Contents []struct {
		Address           ThirdPartyAddress `json:"address"`
		MimeType          string            `json:"mime_type"`
		Name              string            `json:"name,omitempty"`
		Tags              []string          `json:"tags,omitempty"`
		SizeInBytes       string            `json:"size_in_bytes,omitempty"`
		DurationInSeconds string            `json:"duration_in_seconds,omitempty"`
	} `json:"contents,omitempty"`

	Contexts []struct {
		Type     string       `json:"type,omitempty"`
		List     []RSS3ItemID `json:"list,omitempty"`
		ListNext RSS3ListID   `json:"list_next,omitempty"`
	} `json:"@contexts,omitempty"`

	Signature string `json:"signature,omitempty"`
}
