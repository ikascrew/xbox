package xbox

import (
	"fmt"
	"strings"
)

type Event struct {
	Buttons []*Button
	Axes    []*Axis

	error error
}

func (e *Event) push() bool {
	if e.error != nil {
		return true
	}

	if len(e.Buttons) > 0 || len(e.Axes) > 0 {
		return true
	}
	return false
}

func (e *Event) Error() error {
	return e.error
}

func (e *Event) String() string {

	var buf strings.Builder

	for _, elm := range e.Buttons {
		fmt.Fprintf(&buf, "Push [%s]\n", elm.Name)
	}

	for _, elm := range e.Axes {
		fmt.Fprintf(&buf, "Axis [%s][%d]\n", elm.Name, elm.Value)
	}

	return buf.String()
}
