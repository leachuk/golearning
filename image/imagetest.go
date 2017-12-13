package main

import (
	"os"
	"image/gif"
	"image"
	"image/draw"
	"image/png"
	"fmt"
)


func main() {
	file, err := os.Open("1.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	g, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}

	// Define destination boundary.
	r := g.Image[0].Rect
	boundary := image.Rect(0, 0, r.Max.X, r.Max.Y*len(g.Image))
	base := image.NewRGBA(boundary)

	// Copy all Image data into destination boundary sequencially.
	zeroOrigin := image.Point{0, 0}
	for n, img := range g.Image {
		nthBoundary := image.Rect(0, n*r.Max.Y, r.Max.X, (n+1)*r.Max.Y)
		draw.Draw(base, nthBoundary, img, zeroOrigin, draw.Src)
	}

	// Dump image data into file.
	file, err = os.Create("test.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, base)
	if err != nil {
		panic(err)
	}
	fmt.Println("wrote out test.png")
}