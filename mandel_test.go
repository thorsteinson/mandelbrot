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

func TestEscapeIterationCountSmooth(t *testing.T) {
	params := MandelParams{
		MaxIterations: 10000,
		Smooth:        true,
	}

	r1 := EscapeIterationCount(complex(1, 1), params)
	r2 := EscapeIterationCount(complex(100, 5), params)
	r3 := EscapeIterationCount(complex(2, -1), params)
	r4 := EscapeIterationCount(complex(1e6, 1e6), params)
	testVals := []IterationCount{r1, r2, r3, r4}

	for _, r := range testVals {
		if r.Frac < 0 || r.Frac > 1 {
			t.Errorf("Fractional component outside range (0,1): %v", r.Frac)
			t.Logf("%v", r)
		}
	}
}
