package xbox

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"

	"github.com/simulatedsimian/joystick"
	"golang.org/x/xerrors"
)

type Controller struct {
	js      joystick.Joystick
	buttons []bool
	axes    []int

	logger     *log.Logger
	axisMargin int
	duration   int
	rapidFire  bool

	buttonNames []string
	axisNames   []string

	ev   chan *Event
	stop chan bool
}

func defaultController(js joystick.Joystick) *Controller {

	c := Controller{}
	c.js = js
	c.duration = 50

	c.buttons = make([]bool, js.ButtonCount())
	c.axes = make([]int, js.AxisCount())

	c.axisMargin = 2000
	c.rapidFire = false

	return &c
}

func defaultLogger() *log.Logger {
	l := log.New(ioutil.Discard, "[CONT]", log.LstdFlags|log.Lshortfile)
	return l
}

func New(jsId int, opts ...Option) (*Controller, error) {

	js, err := joystick.Open(jsId)
	if err != nil {
		return nil, xerrors.Errorf("Joystick open error: %w", err)
	}

	c := defaultController(js)

	for _, opt := range opts {
		err = opt(c)
		if err != nil {
			return nil, xerrors.Errorf("option set error: %w", err)
		}
	}

	if c.logger == nil {
		c.logger = defaultLogger()
	}

	c.logger.Println(c)

	return c, nil
}

func (c *Controller) TimeDuration() time.Duration {
	return time.Duration(c.duration) * time.Millisecond
}

func (c *Controller) Event() chan *Event {

	if c.ev != nil {
		close(c.ev)
	}
	if c.stop != nil {
		c.stop <- true
	}

	if c.buttonNames == nil || c.axisNames == nil {
		//TODO
	}

	c.ev = make(chan *Event)
	c.stop = make(chan bool)

	go func() {
		ticker := time.NewTicker(c.TimeDuration())
		defer ticker.Stop()
		for doQuit := false; !doQuit; {
			select {
			case <-ticker.C:

				e := Event{}
				err := c.read(&e)
				if err != nil {
					e.error = err
				}

				if e.push() {
					c.ev <- &e
				}
			case <-c.stop:
				close(c.stop)
				c.stop = nil
				return
			}
		}
	}()
	return c.ev
}

func (c *Controller) Terminate() error {
	if c.stop != nil {
		c.stop <- true
	}
	return nil
}

func (c *Controller) Closed() bool {
	if c.stop == nil {
		return true
	}
	return false
}

func (c *Controller) read(e *Event) error {

	js := c.js

	jinfo, err := js.Read()
	if err != nil {
		return fmt.Errorf("Joystick read error[%v]", err)
	}

	e.Buttons = make([]*Button, 0, js.ButtonCount())
	e.Axes = make([]*Axis, 0, js.AxisCount())

	for idx, name := range c.buttonNames {
		if jinfo.Buttons&(1<<uint32(idx)) != 0 {
			ok := true
			if c.buttons[idx] && !c.rapidFire {
				ok = false
			}
			c.buttons[idx] = true

			if ok {
				b := NewButton(idx, name)
				e.Buttons = append(e.Buttons, b)
			}

		} else {
			c.buttons[idx] = false
		}
	}

	for idx, name := range c.axisNames {
		v := jinfo.AxisData[idx]
		c.axes[idx] = v
		if math.Abs(float64(v)) > float64(c.axisMargin) {
			ax := NewAxis(idx, name, v)
			e.Axes = append(e.Axes, ax)
		}
	}

	return nil
}

func (c *Controller) ButtonNames(names ...string) error {

	if c.js == nil {
		return fmt.Errorf("controller initialize error")
	}

	cnt := c.js.ButtonCount()
	if len(names) > cnt {
		return fmt.Errorf("button length error: arguments[%d] > joystick button[%d]", len(names), cnt)
	}
	c.buttonNames = names
	return nil
}

func (c *Controller) AxisNames(names ...string) error {

	if c.js == nil {
		return fmt.Errorf("controller initialize error")
	}

	cnt := c.js.AxisCount()
	if len(names) > cnt {
		return fmt.Errorf("axis length error: arguments[%d] > joystick axis[%d]", len(names), cnt)
	}
	c.axisNames = names
	return nil
}

func (c *Controller) String() string {

	var buf strings.Builder

	fmt.Fprintf(&buf, "\n")
	fmt.Fprintf(&buf, "Joystick Name: %s\n", c.js.Name())
	fmt.Fprintf(&buf, "Button Count : %d\n", c.js.ButtonCount())
	fmt.Fprintf(&buf, "Axis Count   : %d\n", c.js.AxisCount())

	return buf.String()
}
