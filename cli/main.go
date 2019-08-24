package main

import (
	"log"
	"fractal"
)

func main() {
	img := fractal.Render(fractal.DefaultRender)
	err := fractal.ExportPNG(img)
	if err != nil {
		log.Fatalf("Encountered error: %v")
	}
}
