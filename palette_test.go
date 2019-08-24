package fractal

import (
	"testing"
	"image/color"
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

	if len(grays) != res*2 + 1 {
		t.Errorf("Incorrect number of colors found in palette. Expect %v; Found %v", res*2 + 1, len(grays))
	}

	if grays[0] != color.White {
		t.Errorf("Starting color is not white. Found %v", grays[0])
	}

	if grays[len(grays) - 1] != color.White {
		t.Errorf("Ending color is not white. Found %v", grays[len(grays) -1 ])
	}

	if grays[res] != color.Black {
		t.Errorf("Middle color is not black. Found %v", grays[res])
	}
}
