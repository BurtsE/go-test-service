package model

type Statistics struct {
	Received   *uint `json:"received,omitempty"`
	Processed  *uint `json:"processed,omitempty"`
	InProgress *uint `json:"in_progress,omitempty"`
}
