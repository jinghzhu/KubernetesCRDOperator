package events

import (
	"fmt"

	"k8s.io/client-go/tools/cache"

	jinghzhuv1 "github.com/jinghzhu/KubernetesCRD/pkg/crd/jinghzhu/v1"
)

// Event represents the event output by informer.
type Event struct {
	Key         string
	EventType   string
	OldJinghzhu *jinghzhuv1.Jinghzhu
	NewJinghzhu *jinghzhuv1.Jinghzhu
}

// NewEvent returns the pointer to an empty Event.
func NewEvent() *Event {
	return &Event{}
}

// SplitKey returns the namespace and name.
func (e *Event) SplitKey() (string, string, error) {
	return cache.SplitMetaNamespaceKey(e.Key)
}

func (e *Event) String() string {
	return fmt.Sprintf(
		"\tkey = %s\n\teventType = %s\n\told = %+v\n\tnew= %+v",
		e.Key,
		e.EventType,
		e.OldJinghzhu,
		e.NewJinghzhu,
	)
}

const (
	EventAdd    string = "add"
	EventUpdate string = "update"
	EventDelete string = "delete"
)
