package xbox

import (
	"fmt"
	"time"

	"github.com/simulatedsimian/joystick"
)

var duration time.Duration
var handler func(Event) error

func init() {
	SetDuration(100)
	handler = nil
}

func SetDuration(d int) {
	duration = time.Duration(d)
}

func HandleFunc(fn func(Event) error) error {
	handler = fn
	return nil
}

func Open(jsId int) (Event, error) {
	js, err := joystick.Open(jsId)
	if err != nil {
		return Event{}, fmt.Errorf("Joystick open error.[%v]", err)
	}
	e := Event{}
	e.Buttons = make([]bool, js.ButtonCount())
	e.Axes = make([]int, js.AxisCount())

	fmt.Printf("Button:%d\n", js.ButtonCount())
	fmt.Printf("Axes:%d\n", js.AxisCount())

	e.js = js
	return e, nil
}

func Listen(jsId int) error {

	if handler == nil {
		return fmt.Errorf("Call HandleFunc()")
	}

	e, err := Open(jsId)
	if err != nil {
		return fmt.Errorf("Joystick open error.[%v]", err)
	}

	ticker := time.NewTicker(time.Millisecond * duration)
	for doQuit := false; !doQuit; {
		select {
		case <-ticker.C:
			err = read(e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
