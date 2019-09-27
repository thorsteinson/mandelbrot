package main

import (
	"fmt"
	"flag"
)

type Args struct {
	Xres int
	Yres int
}

func ParseArgs() Args {
	xres := flag.Int("xres", 300, "X Resolution")
	yres := flag.Int("yres", 300, "Y Resolution")
	flag.Parse()

	return Args{
		Xres: *xres,
		Yres: *yres,
	}
}

func main() {
	args := ParseArgs()
	fmt.Println("xres:", args.Xres)
	fmt.Println("yres:", args.Yres)
}
