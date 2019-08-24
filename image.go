package fractal

import (
	"fractal/viewport"
	"image"
	"image/color"
	"image/png"
	"os"
)

type IterationImage struct {
	iters []IterationCount
	xres  int
	yres  int
}

type Painter = func(IterationCount) color.Color

// Paint applies a provided coloring function on an IterationImage and
// turns it into a colored image, which can be saved as a PNG or some
// other type of stored bitmap on a machine.
func Paint(p Painter, iterImg IterationImage) *image.RGBA {
	rect := image.Rect(0, 0, iterImg.xres, iterImg.yres)
	img := image.NewRGBA(rect)
	for i, res := range iterImg.iters {
		x := i % iterImg.xres
		y := i / iterImg.xres
		img.Set(x, y, p(res))
	}

	return img
}

func ExportPNG(m *image.RGBA) error {
	err := png.Encode(os.Stdout, m)
	if err != nil {
		return err
	}
	return nil
}

// Render is a workhorse function that generates a complex plane, uses
// the selected algorithm for running the mandelbrot, and colors every
// pixel concurrently
type RenderParams struct {
	VP      viewport.ViewPort
	P       Painter
	Mandel  Algorithm
	MandelP MandelParams
}

var DefaultRender = RenderParams{
	VP:     viewport.New(),
	P:      BlackWhite,
	Mandel: EscapeIterationCount,
	MandelP: MandelParams{
		Bailout:       2,
		MaxIterations: 10000,
	},
}

func Render(params RenderParams) *image.RGBA {
	points := params.VP.Points()
	x, y := params.VP.GetRes()

	iterimg := IterationImage{
		iters: make([]IterationCount, len(points)),
		xres:  x,
		yres:  y,
	}

	// Vast majority of CPU is spent here
	for i, z := range points {
		iterimg.iters[i] = params.Mandel(z, params.MandelP)
	}

	return Paint(params.P, iterimg)
}
