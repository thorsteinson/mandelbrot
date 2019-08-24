package fractal

import (
	"image/color"
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
}
