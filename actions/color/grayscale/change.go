package grayscale

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

type Converter struct {
	DstFile *os.File
}

func (c *Converter) Convert(img image.Image) error {
	bounds := img.Bounds()
	dest := image.NewGray16(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.Gray16Model.Convert(img.At(x, y))
			gray, _ := c.(color.Gray16)
			dest.Set(x, y, gray)
		}
	}
	return jpeg.Encode(c.DstFile, dest, nil)
}
