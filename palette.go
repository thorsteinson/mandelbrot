package fractal

import (
	"image/color"
)

// The simplest possible coloring function. It colors points in the
// Mandelbrot set black, and those not in the set white
func BlackWhite (c IterationCount) color.Gray16 {
	if c.Count < 0 {
		return color.Black
	}
	return color.White
}
