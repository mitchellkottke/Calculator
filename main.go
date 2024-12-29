/*
Calculator calculates basic math expressions.

Usage:

	calculator <expression>

	expression - A complete math expression. Supported operators are:
					+, -, *, /, ^, ()
*/
package main

import (
	"fmt"
	"log"
	"os"

	"calculator/calculations"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := new(app.Window)
		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main() //Boot up the UI window
}

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops
	ans, err := calculations.Evaluate("1+1")
	if err {
		log.Fatal("Calculation failed")
	}

	titleStr := fmt.Sprintf("1 + 1 = %f", ans)
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

// func main() {
// 	var exp string = parseArgs()

// 	ans, failed := calculations.Evaluate(exp)
// 	if !failed {
// 		fmt.Println("Answer is:", ans)
// 	}
// }

// func parseArgs() string {
// 	if len(os.Args) != 2 || len(os.Args[1]) == 0 {
// 		usage(os.Args[0])
// 	}

// 	return strings.Join(os.Args[1:], "")
// }

// func usage(progName string) {
// 	fmt.Println("Usage: ", progName, "<expression>")
// 	fmt.Println("\texpression - The expression to evaluate")
// 	fmt.Println("\t\tSupported operators are: +, -, *, /, ^, ()")
// 	os.Exit(1)
// }
