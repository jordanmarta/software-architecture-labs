package core

type Event string

const (
	EventOpen       Event = "open"
	EventTextChange Event = "text_change"
	EventSave       Event = "save"
)

type Plugin interface {
	Name() string
	OnEvent(event Event, document *Document) Result
}
