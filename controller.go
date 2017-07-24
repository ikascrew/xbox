package xbox

import (
	"fmt"
	"time"

	"github.com/simulatedsimian/joystick"
)

var duration time.Duration
var handler func(Event) error

func init() {
	duration = time.Duration(50)
	handler = nil
}

func SetDuration(d int) {
	duration = time.Duration(d)
}

func HandleFunc(fn func(Event) error) error {
	handler = fn
	return nil
}

func Listen(jsId int) error {

	if handler == nil {
		return fmt.Errorf("Call HandleFunc()")
	}

	js, err := joystick.Open(jsId)
	if err != nil {
		return fmt.Errorf("Joystick open error.[%v]", err)
	}

	e := Event{}
	e.Buttons = make([]bool, js.ButtonCount())
	e.Axes = make([]int, js.AxisCount())

	ticker := time.NewTicker(time.Millisecond * duration)
	for doQuit := false; !doQuit; {
		select {
		case <-ticker.C:
			err = read(js, e)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
