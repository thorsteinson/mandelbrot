package fractal

import (
	"testing"
)

// Tests that the escape iteration count algorithm returns expected
// values for things which are and aren't in the set. This isn't
// exhaustive but helps validate some expectations for a few sample points
func TestEscapeIterationCount(t *testing.T) {
	type testCase struct {
		c        complex128
		isMandel bool
	}

	cases := []testCase{
		{complex(0, 0), true},
		{complex(0.00001, 0.00001), true},
		{complex(-1, 0), true},
		{complex(-1.001, 0), true},
		{complex(1, 1), false},
		{complex(-1, -1), false},
	}

	params := MandelParams{
		Bailout:       2,
		MaxIterations: 1000,
	}

	for _, test := range cases {
		r := EscapeIterationCount(test.c, params)
		if test.isMandel && r.Count >= 0 {
			t.Errorf("Value expected to be in Mandelbrot set: %v. Escaped after %v iterations", test.c, r.Count)
		}
		if !test.isMandel && r.Count < 0 {
			t.Errorf("Value not expected in Mandelbrot set: %v", test.c)
		}
	}
}
