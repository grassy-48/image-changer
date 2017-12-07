package color

import (
	"errors"
	"fmt"
	"image"
	"image-changer/actions/color/grayscale"
	"image-changer/actions/color/negaposi"
	"image-changer/actions/color/sepia"
	"image-changer/global"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const dstFileName string = "%s_%s.jpg"

func Color(c *cli.Context) error {
	arg := c.String("type")

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

	switch arg {
	case "gray":
		gc := grayscale.Converter{
			DstFile: dstfile,
		}
		return gc.Convert(img)
	case "sepia":
		sc := sepia.Converter{
			DstFile: dstfile,
		}
		return sc.Convert(img)
	case "negaposi":
		sc := negaposi.Converter{
			DstFile: dstfile,
		}
		return sc.Convert(img)
	default:
		return errors.New("not support type")
	}
	return nil
}
