package valueobject

// Paging ...
type Paging struct {
	Total int64  `json:"total"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
}
