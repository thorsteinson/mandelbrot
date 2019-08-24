package fractal

import (
	"image/color"
	"math"
)

// The simplest possible coloring function. It colors points in the
// Mandelbrot set black, and those not in the set white
func BlackWhite(c IterationCount) color.Color {
	if c.Count < 0 {
		return color.Black
	}
	return color.White
}

// Generates a grayscale palette to a coloring spectrum which is
// mirrored. Max is the number of maximum interations, for selecting
// the number in the cycle. Finally, resolution is the number of
// grayscale colors to pick for generating the palette.
func Grayscale(resolution int, percentGray float64) color.Palette {
	startingLuminosity := 1.0
	var midpointLuminosity = percentGray * float64(math.MaxUint16)
	deltaLum := (startingLuminosity - midpointLuminosity) / float64(resolution)

	// Construct palette in terms of normalized luminosities
	luminosities := []float64{}
	for i := 0; i <= resolution; i++ {
		lum := startingLuminosity - (float64(i) * deltaLum)
		luminosities = append(luminosities, lum)
	}
	// Mirror the palette
	for i := len(luminosities) - 2; i >= 0; i-- {
		luminosities = append(luminosities, luminosities[i])
	}

	// Convert Palette into actual list of grayscale colors
	palette := make([]color.Color, len(luminosities))
	for i, lum := range luminosities {
		palette[i] = color.Gray16{uint16(lum * float64(math.MaxUint16))}
	}

	return palette
}

// Colors a given index according to the maximum number of iterations,
// and additionally handles linearly interpolating between colors if a
// fractional portion of the interation count is present
func PaintWithPalette(c IterationCount, p color.Palette) color.Color {
	idx := c.Count % len(p)
	color := p[idx]
	if c.Frac > 0 {
		// TODO Linearly interpolate the color in RBG color space
	}
	return color
}
