package model

const (
	InProgress = 1 << iota
	Finished
)

type Message struct {
	Text   string
	Status uint
}
