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
