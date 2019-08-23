package fractal

import (
	"math"
	"testing"
)

// assertRoughEquality will compare the equality of complex numbers,
// and log an Error if they aren't equal as expected
func assertRoughEquality(c1 complex128, c2 complex128, t *testing.T) {
	const epsilon = 0.00000001

	t.Logf("Comparing values %v and %v", c1, c2)
	// test real
	r1 := real(c1)
	r2 := real(c2)
	if math.Abs(r1-r2) > epsilon {
		t.Errorf("Real values differ; expected %v, found %v", r1, r2)
	}

	// test imag
	i1 := imag(c1)
	i2 := imag(c2)
	if math.Abs(i1-i2) > epsilon {
		t.Errorf("Imaginary values differ; expected %v, found %v", i1, i2)
	}
}

func TestSquarePointGeneration(t *testing.T) {
	vp := ViewPort{
		XRes: 2,
		YRes: 2,
		Zoom: 1,
	}

	points := vp.Points()

	testPoints := []complex128{
		complex(-0.5, 0.5),
		complex(0.5, 0.5),
		complex(-0.5, -0.5),
		complex(0.5, -0.5),
	}

	for i := 0; i < len(testPoints); i++ {
		assertRoughEquality(testPoints[i], points[i], t)
	}
}

func TestSquarePointShifted(t *testing.T) {
	vp := ViewPort{
		XRes: 2,
		YRes: 2,
		Zoom: 1,
		C:    complex(10, 10),
	}

	points := vp.Points()

	testPoints := []complex128{
		complex(9.5, 10.5),
		complex(10.5, 10.5),
		complex(9.5, 9.5),
		complex(10.5, 9.5),
	}

	for i := 0; i < len(testPoints); i++ {
		assertRoughEquality(testPoints[i], points[i], t)
	}
}

func TestSquareShiftAndScale(t *testing.T) {
	vp := ViewPort{
		XRes: 4,
		YRes: 4,
		Zoom: 0.5,
		C:    complex(2, 2),
	}

	points := vp.Points()

	testPoints := []complex128{
		// Row 1
		complex(0.5, 3.5),
		complex(1.5, 3.5),
		complex(2.5, 3.5),
		complex(3.5, 3.5),
		// Row 2
		complex(0.5, 2.5),
		complex(1.5, 2.5),
		complex(2.5, 2.5),
		complex(3.5, 2.5),
		// Row 3
		complex(0.5, 1.5),
		complex(1.5, 1.5),
		complex(2.5, 1.5),
		complex(3.5, 1.5),
		// Row 4
		complex(0.5, 0.5),
		complex(1.5, 0.5),
		complex(2.5, 0.5),
		complex(3.5, 0.5),
	}

	for i := 0; i < len(testPoints); i++ {
		assertRoughEquality(testPoints[i], points[i], t)
	}
}

func TestWideViewPort(t *testing.T) {
	vp := ViewPort{
		XRes: 4,
		YRes: 2,
		Zoom: 1,
	}

	points := vp.Points()

	testPoints := []complex128{
		// Row 1
		complex(-1.5, 0.5),
		complex(-0.5, 0.5),
		complex(0.5, 0.5),
		complex(1.5, 0.5),
		// Row 2
		complex(-1.5, -0.5),
		complex(-0.5, -0.5),
		complex(0.5, -0.5),
		complex(1.5, -0.5),
	}

	for i := 0; i < len(testPoints); i++ {
		assertRoughEquality(testPoints[i], points[i], t)
	}
}

func TestTallViewPort(t *testing.T) {
	vp := ViewPort{
		XRes: 2,
		YRes: 4,
		Zoom: 1,
	}

	points := vp.Points()

	testPoints := []complex128{
		// Row 1
		complex(-0.5, 1.5),
		complex(0.5, 1.5),
		// Row 2
		complex(-0.5, 0.5),
		complex(0.5, 0.5),
		// Row 3
		complex(-0.5, -0.5),
		complex(0.5, -0.5),
		// Row 4
		complex(-0.5, -1.5),
		complex(0.5, -1.5),
	}

	for i := 0; i < len(testPoints); i++ {
		assertRoughEquality(testPoints[i], points[i], t)
	}
}
