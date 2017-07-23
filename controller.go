package xbox

import (
	"fmt"
	"math"
	"time"

	"github.com/simulatedsimian/joystick"
)

type Event struct {
	Buttons []bool
	Axes    []int
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
	XBOX
	JOY_L
	JOY_R
)

type Axis int

const (
	JOY_L_HORIZONTAL = Axis(iota)
	JOY_L_VERTICAL
	L2
	JOY_R_HORIZONTAL
	JOY_R_VERTICAL
	R2
	CROSS_HORIZONTAL
	CROSS_VERTICAL
)

//Upper
//Lower
//Left
//Right

func JudgeAxis(e Event, axis Axis) bool {

	if axis == JOY_L_HORIZONTAL || axis == JOY_L_VERTICAL ||
		axis == JOY_R_HORIZONTAL || axis == JOY_R_VERTICAL {
		val := e.Axes[axis]
		if math.Abs(float64(val)) > 6000.0 {
			return true
		}
	} else if axis == L2 || axis == R2 {
		val := e.Axes[axis]
		if val > -30000 {
			return true
		}
	} else if axis == CROSS_HORIZONTAL || axis == CROSS_VERTICAL {
		val := e.Axes[axis]
		if val != 0 {
			return true
		}
	}
	return false
}

func read(js joystick.Joystick, e Event) error {

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

	// TODO error event
	handler(e)

	return nil
}

var handler func(Event) error

func HandlerFunc(fn func(Event) error) error {
	handler = fn
	return nil
}

func Listen(jsId int) error {

	js, err := joystick.Open(jsId)
	if err != nil {
		return fmt.Errorf("Joystick open error.[%v]", err)
	}

	e := Event{}
	e.Buttons = make([]bool, js.ButtonCount())
	e.Axes = make([]int, js.AxisCount())

	ticker := time.NewTicker(time.Millisecond * 100)
	for doQuit := false; !doQuit; {
		select {
		case <-ticker.C:
			err = read(js, e)
			if err != nil {
				//TODO Error Handling
				fmt.Printf("Joystick error[%v]\n", err)
			}
		}
	}
	return nil
}
