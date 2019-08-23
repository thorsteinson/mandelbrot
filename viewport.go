package fractal

type ViewPort struct {
	XRes int
	YRes int
	Zoom float64
	C    complex128
}

// Default is a view port that creates a 500x500 viewing that contains
// the mandelbrot. It's a good starting point that can help to verify
// everything is working as expected.
var Default = ViewPort{
	XRes: 500,
	YRes: 500,
	Zoom: 0.75,
	C:    0.5 + 0i,
}

// Points will return a slice of all the complex numbers that are in
// the provided viewport.
func (vp ViewPort) Points() []complex128 {
	ps := make([]complex128, vp.XRes*vp.YRes)

	topLeft := complex(-1, 1)
	topRight := complex(1, 1)
	normalDist := real(topRight) - real(topLeft)

	var pxLength float64
	if vp.XRes < vp.YRes {
		// Tall image
		// Scale the imaginary part
		topLeft = complex(real(topLeft), float64(vp.YRes)/float64(vp.XRes)*imag(topLeft))
		pxLength = normalDist / float64(vp.XRes)
	} else if vp.XRes > vp.YRes {
		// Wide image
		// Scale the real part
		topLeft = complex(float64(vp.XRes)/float64(vp.YRes)*real(topLeft), imag(topLeft))
		pxLength = normalDist / float64(vp.YRes)
	} else {
		// Square image
		pxLength = normalDist / float64(vp.XRes)
	}

	// Scale the pixel length value
	pxLength /= vp.Zoom

	midpointShift := complex(pxLength/2, -pxLength/2)

	// Scale and shift our reference points
	startingPoint := topLeft/complex(vp.Zoom, 0) + vp.C + midpointShift
	startingX := real(startingPoint)
	p := startingPoint
	for i := 0; i < vp.YRes; i++ {
		for j := 0; j < vp.XRes; j++ {
			idx := i*vp.XRes + j
			ps[idx] = p
			p += complex(pxLength, 0)
		}
		// Reset the real portion and move down
		p = complex(startingX, imag(p))
		p -= complex(0, pxLength)
	}

	return ps
}
