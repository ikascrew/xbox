package main

import (
	"fmt"
	"log"

	"github.com/secondarykey/xbox"
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
	if e.Buttons[xbox.XBOX] {
		return fmt.Errorf("XBOX Button Escape.")
	}
	if e.Buttons[xbox.JOY_R] {
		fmt.Println("Right joystick Button")
	}
	if e.Buttons[xbox.JOY_L] {
		fmt.Println("Left joystick Button")
	}

	if xbox.JudgeAxis(e, xbox.R2) {
		fmt.Printf("R2 Button[%d]\n", e.Axes[xbox.R2])
	}
	if xbox.JudgeAxis(e, xbox.L2) {
		fmt.Printf("L2 Button[%d]\n", e.Axes[xbox.L2])
	}
	if xbox.JudgeAxis(e, xbox.JOY_L_HORIZONTAL) {
		fmt.Printf("Left Joy stick horizontal Button[%d]\n", e.Axes[xbox.JOY_L_HORIZONTAL])
	}
	if xbox.JudgeAxis(e, xbox.JOY_L_VERTICAL) {
		fmt.Printf("Left Joy stick vertical Button[%d]\n", e.Axes[xbox.JOY_L_VERTICAL])
	}
	if xbox.JudgeAxis(e, xbox.JOY_R_HORIZONTAL) {
		fmt.Printf("Right Joy stick horizontal Button[%d]\n", e.Axes[xbox.JOY_R_HORIZONTAL])
	}
	if xbox.JudgeAxis(e, xbox.JOY_R_VERTICAL) {
		fmt.Printf("Right Joy stick vertical Button[%d]\n", e.Axes[xbox.JOY_R_VERTICAL])
	}

	if xbox.JudgeAxis(e, xbox.CROSS_HORIZONTAL) {
		fmt.Printf("Right Joy stick horizontal Button[%d]\n", e.Axes[xbox.CROSS_HORIZONTAL])
	}
	if xbox.JudgeAxis(e, xbox.CROSS_VERTICAL) {
		fmt.Printf("Right Joy stick vertical Button[%d]\n", e.Axes[xbox.CROSS_VERTICAL])
	}

	return nil
}
