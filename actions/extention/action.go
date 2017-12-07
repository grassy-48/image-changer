package extention

import (
	"errors"
	"fmt"
	"image"
	"image-changer/global"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const dstFileName string = "%s_extention.%s"

func Extention(c *cli.Context) error {
	arg := c.String("ext")

	switch arg {
	case "jpeg", "jpg", "gif", "png":
	default:
		return errors.New("not support extention")
	}
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
	case "jpeg", "jpg":
		err = jpeg.Encode(dstfile, img, nil)
	case "gif":
		err = gif.Encode(dstfile, img, nil)
	case "png":
		err = png.Encode(dstfile, img)
	}
	return err
}
