package processstatuscode

//go:generate stringer -type=Status
type Status int64

// Process status code
const (
	Success       Status = 200
	Drop          Status = 500
	Retry         Status = 400
	FailReproduce Status = 302
)
