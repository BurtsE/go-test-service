package model

const (
	Created = 1 << iota
	InProgress
	Finished
)

type Message struct {
	UUID   uint64 `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Status uint   `json:"status"`
}
