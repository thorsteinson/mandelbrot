package viewport

import (
	"image"
)

type vp struct {
	xres int
	yres int
	zoom float64
	c    complex128
}

type View interface {
	Points() []complex128
	Rect() image.Rectangle
}

// Points will return a slice of all the complex numbers that are in
// the provided viewport.
func (vp vp) Points() []complex128 {
	ps := make([]complex128, vp.xres*vp.yres)

	topLeft := complex(-1, 1)
	topRight := complex(1, 1)
	normalDist := real(topRight) - real(topLeft)

	var pxLength float64
	if vp.xres < vp.yres {
		// Tall image
		// Scale the imaginary part
		topLeft = complex(real(topLeft), float64(vp.yres)/float64(vp.xres)*imag(topLeft))
		pxLength = normalDist / float64(vp.xres)
	} else if vp.xres > vp.yres {
		// Wide image
		// Scale the real part
		topLeft = complex(float64(vp.xres)/float64(vp.yres)*real(topLeft), imag(topLeft))
		pxLength = normalDist / float64(vp.yres)
	} else {
		// Square image
		pxLength = normalDist / float64(vp.xres)
	}

	// Scale the pixel length value
	pxLength /= vp.zoom

	midpointShift := complex(pxLength/2, -pxLength/2)

	// Scale and shift our reference points
	startingPoint := topLeft/complex(vp.zoom, 0) + vp.c + midpointShift
	startingX := real(startingPoint)
	p := startingPoint
	for i := 0; i < vp.yres; i++ {
		for j := 0; j < vp.xres; j++ {
			idx := i*vp.xres + j
			ps[idx] = p
			p += complex(pxLength, 0)
		}
		// Reset the real portion and move down
		p = complex(startingX, imag(p))
		p -= complex(0, pxLength)
	}

	return ps
}

// New returns a vp. It's the only public way to get a
// viewport. By default it returns a default view that includes the
// classic picture most associate with the Mandelbrot set
func New() vp {
	return vp{
		xres: 500,
		yres: 500,
		zoom: 0.75,
		c:    -0.5 + 0i,
	}
}

// Res safely sets the resolution for the viewport. Panics when passed
// values less than 1.
func (vp vp) Res(x int, y int) vp {
	if x < 1 || y < 1 {
		panic("Attempted to set resolution of View Port below 1")
	}

	vp.xres = x
	vp.yres = y

	return vp
}

// Zoom safely sets the zoom level of the viewport. If a negative zoom
// level is provided the program will crash.
func (vp vp) Zoom(z float64) vp {
	if z < 0 {
		panic("Attempted to set negative zoom level on View Port")
	}

	vp.zoom = z

	return vp
}

// Rect returns a rectangle generated from the provided view port
// which can be passed to certain image functions
func (vp vp) Rect() image.Rectangle {
	return image.Rect(0, 0, vp.xres, vp.yres)
}
