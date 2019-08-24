package viewport

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

func TestPointGeneration(t *testing.T) {
	type PointCompTest struct {
		name     string
		expected []complex128
		viewport ViewPort
	}

	tests := []PointCompTest{
		{"Square Point Generation",
			[]complex128{
				complex(-0.5, 0.5),
				complex(0.5, 0.5),
				complex(-0.5, -0.5),
				complex(0.5, -0.5),
			},
			ViewPort{
				xres: 2,
				yres: 2,
				zoom: 1,
			},
		},

		{"Square Point Shifted",
			[]complex128{
				complex(9.5, 10.5),
				complex(10.5, 10.5),
				complex(9.5, 9.5),
				complex(10.5, 9.5),
			},
			ViewPort{
				xres: 2,
				yres: 2,
				zoom: 1,
				c:    complex(10, 10),
			},
		},

		{"Square Shift and Scale",
			[]complex128{
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
			},
			ViewPort{
				xres: 4,
				yres: 4,
				zoom: 0.5,
				c:    complex(2, 2),
			},
		},
		{"Wide View Port",
			[]complex128{
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
			},
			ViewPort{
				xres: 4,
				yres: 2,
				zoom: 1,
			},
		},
		{"Tall View Port",
			[]complex128{
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
			},
			ViewPort{
				xres: 2,
				yres: 4,
				zoom: 1,
			},
		},
	}

	for _, test := range tests {
		t.Logf("Testing: %v", test.name)
		t.Logf("View Port: %v", test.viewport)
		points := test.viewport.Points()
		for i, p := range points {
			assertRoughEquality(p, points[i], t)
		}
	}
}

// Helper function for testing panics
func resPanics(x int, y int) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	vp := New()
	vp.Res(x, y)
	return panicked
}

func TestResolutionModification(t *testing.T) {
	vp := ViewPort{}

	vp = vp.Res(100, 200)

	if vp.xres != 100 {
		t.Error("Failed to set the X resolution")
	}
	if vp.yres != 200 {
		t.Error("Failed to set the Y resolution")
	}

	if !resPanics(0, 10) || !resPanics(10, 0) {
		t.Error("Setting resolution to 0 did not panic")
	}

	if !resPanics(-1, 100) || !resPanics(100, -1) {
		t.Error("Setting resolution below 0 did not panic")
	}
}

func zoomPanics(z float64) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	vp := New()
	vp.Zoom(z)
	return panicked
}

func TestZoomModification(t *testing.T) {
	vp := ViewPort{}

	vp = vp.Zoom(1000)

	if vp.zoom != 1000 {
		t.Error("Failed to set the zoom level")
	}

	if !zoomPanics(-1) {
		t.Error("Setting negative zoom did not panic")
	}
}
