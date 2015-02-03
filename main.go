package main

import (
	"code.google.com/p/draw2d/draw2d"

	"image"
)

func main() {
	// This is an example of drawing a square with draw2d.
	// Your Logo implementation will need to parse a text file, understand each command
	// and then execute it.  You'll probably need to keep state of your position and heading at least.
	// The turtle usually starts in the centre of the screen.
	screenWidth, screenHeight := 400, 400
	i := image.NewRGBA(image.Rect(0, 0, screenWidth, screenHeight))
	homeX, homeY := float64(screenWidth) / 2.0, float64(screenHeight) / 2.0

	gc := draw2d.NewGraphicContext(i)
	gc.MoveTo(homeX, homeY)
	gc.LineTo(homeX + 100.0, homeY)
	gc.Stroke()
	gc.MoveTo(homeX + 100.0, homeY)
	gc.LineTo(homeX + 100.0, homeY - 100.0)
	gc.Stroke()
	gc.MoveTo(homeX + 100.0, homeY - 100.0)
	gc.LineTo(homeX, homeY - 100.0)
	gc.Stroke()
	gc.MoveTo(homeX, homeY - 100.0)
	gc.LineTo(homeX, homeY)
	gc.Stroke()
	saveToPngFile("Logo-output.png", i)
}
