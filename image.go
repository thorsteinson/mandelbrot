package fractal

import (
	"fractal/viewport"
	"image"
	"image/color"
	"image/png"
	"os"
	"runtime"
	"sync"
)

type IterationImage struct {
	iters []IterationCount
	rect  image.Rectangle
}

type Painter = func(IterationCount) color.Color

// Paint applies a provided coloring function on an IterationImage and
// turns it into a colored image, which can be saved as a PNG or some
// other type of stored bitmap on a machine.
func Paint(p Painter, iterImg IterationImage) *image.RGBA {
	img := image.NewRGBA(iterImg.rect)
	for i, res := range iterImg.iters {
		x := i % iterImg.rect.Dx()
		y := i / iterImg.rect.Dy()
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
	VP      viewport.View
	P       Painter
	Mandel  Algorithm
	MandelP MandelParams
}

var DefaultRender = RenderParams{
	VP:     viewport.New(),
	P:      BlackWhite,
	Mandel: EscapeIterationCount,
	MandelP: MandelParams{
		MaxIterations: 10000,
	},
}

func Render(params RenderParams) *image.RGBA {
	type work struct {
		index int
		point complex128
	}

	points := params.VP.Points()

	iterimg := IterationImage{
		iters: make([]IterationCount, len(points)),
		rect:  params.VP.Rect(),
	}

	workers := runtime.NumCPU()

	inC := make(chan work)

	wg := sync.WaitGroup{}
	wg.Add(workers)

	// Set workers to handle input from the channel
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for w := range inC {
				iterimg.iters[w.index] = params.Mandel(w.point, params.MandelP)
			}
		}()
	}

	// Send points we want to process to the workers
	for i, z := range points {
		inC <- work{i, z}
	}
	close(inC)

	wg.Wait()

	return Paint(params.P, iterimg)
}
