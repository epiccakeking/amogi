package main

import (
	"fmt"
	"image"
	"os"

	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
)

var amogs = [4][][2]int{
	{
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 1},
		{1, 1},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
		{1, 3},
		{3, 3},
	},
	{
		{0, 0},
		{2, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{2, 2},
		{3, 2},
		{0, 3},
		{1, 3},
		{2, 3},
	},
}

func drawAmog(imageOut *image.RGBA, col color.Color, offsets [][2]int, x, y int) {
	for _, offset := range offsets {
		imageOut.Set(x+offset[0], y+offset[1], col)
	}
}

func main() {
	var filepath string
	fmt.Scanln(&filepath)
	fmt.Printf("\"%s\"", filepath)
	reader, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	imageIn, _, err := image.Decode(reader)
	if err != nil {
		panic(err)
	}
	// Determine size ouf output
	inSize := imageIn.Bounds().Size()
	// Amogus calculations are too hard
	bounds := image.Rect(
		-inSize.Y/2, 0,
		inSize.X*7/2+1, (inSize.Y*7+inSize.X%2)/2+inSize.X/2*3+inSize.X%2,
	)

	imageOut := image.NewRGBA(bounds)
	// Finally, we can draw the amogi
	for x := 0; x < inSize.X; x++ {
		for y := 0; y < inSize.Y; y++ {
			yOut := (y*7+x%2)/2 + x/2*3 + x%2
			drawAmog(imageOut, imageIn.At(x, y), amogs[(x+y)%2], (x*7+1)/2-(y+x%2)/2, yOut)
		}
	}
	writer, err := os.Create("output.png")
	defer writer.Close()
	if err != nil {
		panic(err)
	}
	png.Encode(writer, imageOut)
}
