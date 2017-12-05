package sepia

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Converter struct {
	DstFile *os.File
}

func (c *Converter) Convert(img image.Image) error {
	bounds := img.Bounds()
	dest := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.NRGBAModel.Convert(img.At(x, y))
			col := c.(color.NRGBA)
			avg := (float32(col.R) + float32(col.G) + float32(col.B)) / 3
			cg := float32(col.G) * 0.8
			cb := float32(col.B) * 0.55
			dest.Set(x, y, color.NRGBA{uint8(avg), uint8(cg), uint8(cb), col.A})
		}
	}
	png.Encode(c.DstFile, dest)
	return nil
}
