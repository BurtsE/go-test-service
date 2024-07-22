package model

const (
	InProgress = 1 << iota
	Finished
)

type Message struct {
	Text   string `json:"text,omitempty"`
	Status uint   `json:"status"`
}
