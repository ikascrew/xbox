package xbox

import (
	"fmt"

	"github.com/simulatedsimian/joystick"
)

type Event struct {
	Buttons []bool
	Axes    []int
	js      joystick.Joystick
}

type Button int

const (
	A = Button(iota)
	B
	X
	Y
	L1
	R1
	BACK
	START
)

type Axis int

const (
	CROSS_HORIZONTAL = Axis(iota)
	CROSS_VERTICAL
)

func JudgeAxis(e Event, axis Axis) bool {
	if e.Axes[axis] != 0 {
		return true
	}
	return false
}

func (e Event) GetEvent() (Event, error) {

	jinfo, err := e.js.Read()
	if err != nil {
		return Event{}, err
	}

	for button := 0; button < e.js.ButtonCount(); button++ {
		if jinfo.Buttons&(1<<uint32(button)) != 0 {
			e.Buttons[button] = true
		} else {
			e.Buttons[button] = false
		}
	}

	for axis := 0; axis < e.js.AxisCount(); axis++ {
		e.Axes[axis] = jinfo.AxisData[axis]
	}
	return e, nil
}

func read(e Event) error {

	js := e.js

	jinfo, err := js.Read()
	if err != nil {
		return fmt.Errorf("Joystick read error[%v]", err)
	}

	for button := 0; button < js.ButtonCount(); button++ {
		if jinfo.Buttons&(1<<uint32(button)) != 0 {
			e.Buttons[button] = true
		} else {
			e.Buttons[button] = false
		}
	}

	for axis := 0; axis < js.AxisCount(); axis++ {
		e.Axes[axis] = jinfo.AxisData[axis]
	}

	return handler(e)
}
