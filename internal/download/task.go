package download

import "time"

type Status string

const (
	Pending     Status = "pending"
	Downloading Status = "downloading"
	Completed   Status = "completed"
	Failed      Status = "failed"
)

type Task struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Output    string    `json:"output"`
	Status    Status    `json:"status"`
	Bytes     int64     `json:"bytes"`
	Total     int64     `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
