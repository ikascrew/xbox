package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ikascrew/xbox"
)

func main() {

	err := run()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fmt.Println("success")
}

func run() error {

	c, err := createController(0)
	if err != nil {
		return err
	}

	ch := c.Event()
	for {
		select {
		case ev := <-ch:
			if ev.Error() != nil {
				fmt.Printf("%+v\n", ev.Error())
				c.Terminate()
			} else {
				fmt.Print(ev)
			}
		default:
		}

		if c.Closed() {
			break
		}
	}

	return nil
}

const (
	Button10Axis7 = true
	Button8Axis2  = false
)

func createController(id int) (*xbox.Controller, error) {

	c, err := xbox.New(id,
		xbox.Logger(log.New(os.Stdout, "[XBOX]", log.LstdFlags|log.Lshortfile)),
		xbox.Duration(40),
		xbox.AxisMargin(3000),
	)
	if err != nil {
		return nil, err
	}

	if Button10Axis7 {
		err = c.ButtonNames("A", "B", "X", "Y", "L", "R", "BACK", "START", "L_JOY", "R_JOY")
		if err != nil {
			return nil, err
		}

		err = c.AxisNames("LEFT_JOY_H", "LEFT_JOY_V", "ZLR", "RIGHT_JOY_V", "RIGHT_JOY_H", "CROSS_H", "CROSS_V")
		if err != nil {
			return nil, err
		}
	} else if Button8Axis2 {
		err = c.ButtonNames("A", "B", "X", "Y", "L", "R", "BACK", "START")
		if err != nil {
			return nil, err
		}
		err = c.AxisNames("CROSS_H", "CROSS_V")
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
