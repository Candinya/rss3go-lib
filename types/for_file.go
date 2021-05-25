package types

// These types are ONLY for storage use!!
// For only signature-oriented types, please check for_sign.go

// ** IMPORTANT: Please obey the official JS/TS type for strict order of types' properties.

// RSS3Base4F Common attributes for each file
type RSS3Base4F struct {
	//Id          interface{} `json:"id"` // RSS3ID | RSS3ItemsID | RSS3ListID
	Version     string `json:"@version"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	Signature   string `json:"signature,omitempty"`
}

// RSS34F file, Entrance
type RSS34F struct {
	Id RSS3ID `json:"id"`

	*RSS3Base4F

	Profile *RSS3Profile4F `json:"profile"`

	Items     []RSS3Item4F `json:"items"`
	ItemsNext RSS3ItemsID  `json:"items_next,omitempty"`

	Links []RSS3Links4F `json:"links,omitempty"`

	Backlinks []RSS3Backlinks4F `json:"@backlinks,omitempty"`

	Assets []RSS3Assets4F `json:"assets,omitempty"`
}

type RSS3Profile4F struct {
	Name      string            `json:"name,omitempty"`
	Avatar    ThirdPartyAddress `json:"avatar,omitempty"`
	Bio       string            `json:"bio,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
	Signature string            `json:"signature,omitempty"`
}

type RSS3Links4F struct {
	Type      string     `json:"type"`
	Tags      []string   `json:"tags,omitempty"`
	List      []RSS3ID   `json:"list"`
	ListNext  RSS3ListID `json:"list_next,omitempty"`
	Signature string     `json:"signature,omitempty"`
}

type RSS3Backlinks4F struct {
	Type     string     `json:"type"`
	List     []RSS3ID   `json:"list"`
	ListNext RSS3ListID `json:"list_next,omitempty"`
}

type RSS3Assets4F struct {
	Type    string   `json:"type"`
	Tags    []string `json:"tags,omitempty"`
	Content string   `json:"content"`
}

// RSS3Items4F file
type RSS3Items4F struct {
	Id RSS3ItemsID `json:"id"`

	*RSS3Base4F

	Items     []RSS3Item4F `json:"items"`
	ItemsNext RSS3ItemsID  `json:"items_next,omitempty"`
}

// RSS3List4F file
type RSS3List4F struct {
	Id RSS3ListID `json:"id"`

	*RSS3Base4F

	List []interface{} `json:"list"` // [] RSS3ID | [] RSS3ItemID

	ListNext RSS3ListID `json:"list_next,omitempty"`
}

// RSS3Item4F file
type RSS3Item4F struct {
	Authors       []RSS3ID `json:"authors,omitempty"`
	Title         string   `json:"title,omitempty"`
	Id            string   `json:"id"`
	Summary       string   `json:"summary,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	DatePublished string   `json:"date_published,omitempty"`
	DateModified  string   `json:"date_modified,omitempty"`

	Type     string     `json:"type,omitempty"`
	Upstream RSS3ItemID `json:"upstream,omitempty"`

	Contents []RSS3ItemContent4F `json:"contents,omitempty"`

	Contexts []RSS3ItemContext4F `json:"@contexts,omitempty"`

	Signature string `json:"signature,omitempty"`
}

type RSS3ItemContent4F struct {
	Address           ThirdPartyAddress `json:"address"`
	MimeType          string            `json:"mime_type"`
	Name              string            `json:"name,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	SizeInBytes       string            `json:"size_in_bytes,omitempty"`
	DurationInSeconds string            `json:"duration_in_seconds,omitempty"`
}

type RSS3ItemContext4F struct {
	Type     string       `json:"type,omitempty"`
	List     []RSS3ItemID `json:"list,omitempty"`
	ListNext RSS3ListID   `json:"list_next,omitempty"`
}
