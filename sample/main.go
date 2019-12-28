package main

import (
	"fmt"
	"log"

	"github.com/ikascrew/xbox"
)

func main() {
	xbox.HandleFunc(printButton)
	err := xbox.Listen(0)
	if err != nil {
		log.Fatal(err)
	}
}

func printButton(e xbox.Event) error {

	if e.Buttons[xbox.A] {
		fmt.Println("A Button")
	}
	if e.Buttons[xbox.B] {
		fmt.Println("B Button")
	}
	if e.Buttons[xbox.Y] {
		fmt.Println("Y Button")
	}
	if e.Buttons[xbox.X] {
		fmt.Println("X Button")
	}
	if e.Buttons[xbox.L1] {
		fmt.Println("L1 Button")
	}
	if e.Buttons[xbox.R1] {
		fmt.Println("R1 Button")
	}
	if e.Buttons[xbox.BACK] {
		fmt.Println("Back Button")
	}
	if e.Buttons[xbox.START] {
		fmt.Println("START Button")
	}

	if xbox.JudgeAxis(e, xbox.CROSS_HORIZONTAL) {
		fmt.Printf("cross button horizontal Button[%d]\n", e.Axes[xbox.CROSS_HORIZONTAL])
	}
	if xbox.JudgeAxis(e, xbox.CROSS_VERTICAL) {
		fmt.Printf("cross button vertical Button[%d]\n", e.Axes[xbox.CROSS_VERTICAL])
	}

	return nil
}
