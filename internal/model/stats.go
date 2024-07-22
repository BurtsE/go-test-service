package model

type Statistics struct {
	Processed   uint `json:"processed,omitempty"`
	UnProcessed uint `json:"in_progress,omitempty"`
}
