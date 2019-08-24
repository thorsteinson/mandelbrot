package fractal

import (
	"math/cmplx"
)

// IterationCount is returned by coloring functions. More advanced
// ones will return a fractional portion. If the count returned is
// negative, then that should be interpreted as a point being in the
// set.
type IterationCount struct {
	Count int
	Frac  float64
}

type MandelParams struct {
	Bailout       float64
	MaxIterations int
}

type Algorithm = func (complex128, MandelParams) IterationCount

// EscapeIterationCount applies a simple algorithm to get an integer
// color for when a point escapes.
func EscapeIterationCount(c complex128, p MandelParams) IterationCount {

	// Move the panic checks to slice production, so we don't do these
	// for each pixel and just do it when the params are first created
	if p.MaxIterations < 1 {
		panic("Max Iterations cannot be less than 1`")
	}
	if p.Bailout < 0 {
		panic("Bailout radius cannot be negative")
	}

	var z complex128
	i := 0
	for i < p.MaxIterations && cmplx.Abs(z) < p.Bailout {
		z = z*z + c
		i++
	}

	if i == p.MaxIterations {
		return IterationCount{Count: -1}
	}

	return IterationCount{
		Count: i,
	}
}
