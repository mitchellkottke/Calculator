/*
Calculator calculates basic math expressions.

Usage:

	calculator <expression>

	expression - A complete math expression. Supported operators are:
					+, -, *, /, ^, ()
*/

// Gio docs: https://pkg.go.dev/gioui.org/widget#pkg-overview

package main

import (
	"fmt"
	"log"
	"os"

	"calculator/calculations"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

/*
	Button Layout:
	|  %  |  C  |     <     |
	|  (  |  )  |  ^  |  /  |
	|  7  |  8  |  9  |  *  |
	|  4  |  5  |  6  |  -  |
	|  1  |  2  |  3  |  +  |
	|     0     |  .  |  =  |
*/

const ( //Button position
	MOD = iota
	CLEAR
	BACK
	LPAREN
	RPAREN
	EXP
	DIV
	SEVEN
	EIGHT
	NINE
	MULT
	FOUR
	FIVE
	SIX
	SUB
	ONE
	TWO
	THREE
	ADD
	ZERO
	DOT
	EQ
)

type button struct {
	btn    widget.Clickable //Button object
	value  int              //Button value
	isWide bool             //If button should be drawn as a double button
}

type state struct {
	buttons []button
	outText string
}

var progState state

func main() {
	go processing_thread() //Start processing thread

	app.Main() //Boot up the UI window
}

func processing_thread() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	window := new(app.Window)
	theme := material.NewTheme()
	var ops op.Ops

	ans, err := calculations.Evaluate("1+1")
	if err {
		log.Fatal("Calculation failed")
	}

	titleStr := fmt.Sprintf("1 + 1 = %g", ans)
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent: //Window is closed
			return e.Err
		case app.FrameEvent: //Render cycle
			//Graphics context
			gtx := app.NewContext(&ops, e)

			//Make title label
			title := material.H1(theme, titleStr)

			//Draw the title
			title.Layout(gtx)

			//Render
			e.Frame(gtx.Ops)
		}
	}
}

func stateInit(s *state) {
	s.outText = ""
	// s.buttons = {
	// 	{}
	// }

	//s.buttons[0].btn.Clicked()
}
