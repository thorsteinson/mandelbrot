package fractal

import (
	"image/color"
	"math"
	"testing"
)

// Tests that we get black colors for points in the set
func TestBlackWhite(t *testing.T) {

	if c := BlackWhite(IterationCount{Count: -1}); c != color.Black {
		t.Error("Failed to color result black")
	}
	if c := BlackWhite(IterationCount{Count: 100}); c != color.White {
		t.Error("Failed to color result white")
	}
}

// Tests that grayscale generates a palette as expected
func TestGrayscale(t *testing.T) {

	res := 5
	// Construct a grayscale palette that varies between complete
	// white and black, with 5 colors mirrored
	grays := Grayscale(res, 0)

	if len(grays) != res*2+1 {
		t.Errorf("Incorrect number of colors found in palette. Expect %v; Found %v", res*2+1, len(grays))
	}

	if grays[0] != color.White {
		t.Errorf("Starting color is not white. Found %v", grays[0])
	}

	if grays[len(grays)-1] != color.White {
		t.Errorf("Ending color is not white. Found %v", grays[len(grays)-1])
	}

	if grays[res] != color.Black {
		t.Errorf("Middle color is not black. Found %v", grays[res])
	}
}

func TestPaintWithPalette(t *testing.T) {
	c0 := color.Gray16{1}
	c1 := color.Gray16{2}
	c2 := color.Gray16{3}
	c3 := color.Gray16{4}
	palette := []color.Color{c0, c1, c2, c3}

	var c IterationCount
	var lookup color.Color

	c = IterationCount{Count: 0}
	lookup = PaintWithPalette(c, palette)
	if lookup != c0 {
		t.Errorf("Failed to lookup proper color. Expected %v, Found %v", c0, lookup)
	}

	c = IterationCount{Count: 1}
	lookup = PaintWithPalette(c, palette)
	if lookup != c1 {
		t.Errorf("Failed to lookup proper color. Expected %v, Found %v", c1, lookup)
	}

	// Test cycling through the palette
	c = IterationCount{Count: 4}
	lookup = PaintWithPalette(c, palette)
	if lookup != c0 {
		t.Errorf("Failed to lookup proper color. Expected %v, Found %v", c0, lookup)
	}

	// Test points in the Mandelbrot set themselves
	c = IterationCount{Count: -1}
	lookup = PaintWithPalette(c, palette)
	if lookup != color.Black {
		t.Errorf("Negative iteration failed to resolve as black color")
	}

}

func testColors(expectedClr color.Color, clr color.Color, t *testing.T) {
	rExpected, gExpected, bExpected, aExpected := expectedClr.RGBA()
	r, g, b, a := clr.RGBA()
	if r != rExpected {
		t.Errorf("Red color channel incorrectly interpolated. Expected %v, Found %v", rExpected, r)
	}
	if g != gExpected {
		t.Errorf("Green color channel incorrectly interpolated. Expected %v, Found %v", gExpected, g)
	}
	if b != bExpected {
		t.Errorf("Blue color channel incorrectly interpolated. Expected %v, Found %v", bExpected, b)
	}
	if a != aExpected {
		t.Errorf("Alpha color channel incorrectly interpolated. Expected %v, Found %v", aExpected, a)
	}
}

func TestPaintWithPaletteInterpolation(t *testing.T) {
	var palette []color.Color
	palette = []color.Color{color.Black, color.White}

	var c IterationCount
	var lookup color.Color

	// Do a simple grayscale test
	c = IterationCount{Count: 0, Frac: 0.5}
	lookup = PaintWithPalette(c, palette)
	// We use Gray instead of Gray16 because there's a loss of
	// precision as we go from 16 to 8 bit.
	testColors(color.Gray{math.MaxUint8 / 2}, lookup, t)

	c = IterationCount{Count: 0, Frac: 0.25}
	lookup = PaintWithPalette(c, palette)
	testColors(color.Gray{math.MaxUint8 / 4}, lookup, t)

	// Test the palette in the reverse direction
	c = IterationCount{Count: 0, Frac: 0.5}
	palette = []color.Color{color.White, color.Black}
	lookup = PaintWithPalette(c, palette)
	testColors(color.Gray{math.MaxUint8 / 2}, lookup, t)

	// Test with just the red channel. We should get 50% red
	c = IterationCount{Count: 0, Frac: 0.5}
	pureRed := color.RGBA{math.MaxUint8, 0, 0, math.MaxUint8}
	palette = []color.Color{color.Black, pureRed}
	lookup = PaintWithPalette(c, palette)
	testColors(color.RGBA{math.MaxUint8 / 2, 0, 0, math.MaxUint8}, lookup, t)

	// Test with just the green channel. We should get 50% green
	pureGreen := color.RGBA{0, math.MaxUint8, 0, math.MaxUint8}
	palette = []color.Color{color.Black, pureGreen}
	lookup = PaintWithPalette(c, palette)
	testColors(color.RGBA{0, math.MaxUint8 / 2, 0, math.MaxUint8}, lookup, t)

	// Test with just the blue channel. We should get 50% blue
	pureBlue := color.RGBA{0, 0, math.MaxUint8, math.MaxUint8}
	palette = []color.Color{color.Black, pureBlue}
	lookup = PaintWithPalette(c, palette)
	testColors(color.RGBA{0, 0, math.MaxUint8 / 2, math.MaxUint8}, lookup, t)
}
