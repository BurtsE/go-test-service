package model

type Status uint
const (
	StatusUndefined Status = iota
	Created
	InProgress
	Finished
)

type Message struct {
	UUID   uint64 `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Status Status   `json:"status"`
}
