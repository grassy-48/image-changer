package mosaic

import (
	"fmt"
	"image"
	"image-changer/global"
	"image/color"
	"image/jpeg"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const dstFileName string = "%s_mosaic_%d.png"

func Mosaic(c *cli.Context) error {
	arg := c.Int("particle")

	file, err := os.Open(fmt.Sprintf("%s%s.%s",
		global.SrcImagePath,
		global.SrcImageFileName,
		global.SrcImageFileExtention,
	))
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dstfile, err := os.Create(
		global.DstImagePath + fmt.Sprintf(
			dstFileName,
			global.SrcImageFileName,
			arg,
		),
	)
	if err != nil {
		return err
	}
	defer dstfile.Close()
	bounds := img.Bounds()
	dest := image.NewRGBA(bounds)

	for y := bounds.Min.Y + (arg-1)/2; y < bounds.Max.Y; y = y + arg {
		for x := bounds.Min.X + (arg-1)/2; x < bounds.Max.X; x = x + arg {
			var cr, cg, cb float32
			var alpha uint8
			for j := y - (arg-1)/2; j <= y+(arg-1)/2; j++ {
				for i := x - (arg-1)/2; i <= x+(arg-1)/2; i++ {
					if i >= 0 && j >= 0 && i < bounds.Max.X && j < bounds.Max.Y {
						c := color.RGBAModel.Convert(img.At(i, j))
						col := c.(color.RGBA)
						cr += float32(col.R)
						cg += float32(col.G)
						cb += float32(col.B)
						alpha = col.A
					}
				}
			}
			cr = cr / float32(arg*arg)
			cg = cg / float32(arg*arg)
			cb = cb / float32(arg*arg)
			for j := y - (arg-1)/2; j <= y+(arg-1)/2; j++ {
				for i := x - (arg-1)/2; i <= x+(arg-1)/2; i++ {
					if i >= 0 && j >= 0 && i < bounds.Max.X && j < bounds.Max.Y {
						dest.Set(i, j, color.RGBA{uint8(cr), uint8(cg), uint8(cb), alpha})
					}
				}
			}
		}
	}
	return jpeg.Encode(dstfile, dest, nil)
}
