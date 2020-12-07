package events

import (
	"fmt"

	jinghzhuv1 "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1"
)

// Event represents the event output by informer.
type Event struct {
	Key         string
	EventType   string
	Namespace   string
	OldJinghzhu *jinghzhuv1.Jinghzhu
	NewJinghzhu *jinghzhuv1.Jinghzhu
}

// NewEvent returns the pointer to an empty Event.
func NewEvent() *Event {
	return &Event{}
}

func (e *Event) String() string {
	return fmt.Sprintf(
		"\tkey = %s\n\teventType = %s\n\tnamespace = %s\n\toldJinghzhu = %+v\n\tnewJinghzhu= %+v",
		e.Key,
		e.EventType,
		e.Namespace,
		e.OldJinghzhu,
		e.NewJinghzhu,
	)
}

const (
	EventAdd    string = "add"
	EventUpdate string = "update"
	EventDelete string = "delete"
)
