package fractal

import (
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
