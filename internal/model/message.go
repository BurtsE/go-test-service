package model

var (
	Created    EnumStatus = &created{}
	InProgress EnumStatus = &inProgress{}
	Finished   EnumStatus = &finished{}
)

type EnumStatus interface{ isEnumRole() }

type created struct{ EnumStatus }
type inProgress struct{ EnumStatus }
type finished struct{ EnumStatus }

type Message struct {
	UUID   uint64     `json:"id,omitempty"`
	Text   string     `json:"text,omitempty"`
	Status EnumStatus `json:"status"`
}
