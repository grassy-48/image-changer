package negaposi

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
			cr := uint8(255 - int(col.R))
			cg := uint8(255 - int(col.G))
			cb := uint8(255 - int(col.B))
			dest.Set(x, y, color.NRGBA{cr, cg, cb, col.A})
		}
	}
	png.Encode(c.DstFile, dest)
	return nil
}
