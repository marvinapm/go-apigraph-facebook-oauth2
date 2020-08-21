package repository_model

type ResponseBody struct {
	Data   []Datum `json:"data"`
	Paging Paging  `json:"paging"`
}

type Datum struct {
	ID        string `json:"id"`
	Caption   string `json:"caption"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
	Username  string `json:"username"`
	Owner     Owner  `json:"owner"`
	Permalink string `json:"permalink"`
	Timestamp string `json:"timestamp"`
}

type Owner struct {
	ID string `json:"id"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
