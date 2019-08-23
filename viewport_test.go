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
				XRes: 2,
				YRes: 2,
				Zoom: 1,
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
				XRes: 2,
				YRes: 2,
				Zoom: 1,
				C:    complex(10, 10),
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
				XRes: 4,
				YRes: 4,
				Zoom: 0.5,
				C:    complex(2, 2),
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
				XRes: 4,
				YRes: 2,
				Zoom: 1,
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
				XRes: 2,
				YRes: 4,
				Zoom: 1,
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
