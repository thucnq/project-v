package paginator

type Cursor struct {
	NextPageToken *string `json:"next_page_token" query:"next_page_token"`
	PrevPageToken *string `json:"prev_page_token" query:"prev_page_token"`
}
