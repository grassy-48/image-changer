package origin

import (
	"fmt"
	"image-changer/global"
	"io"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const dstFileName string = "%s_origin.%s"

func Origin(c *cli.Context) error {
	file, err := os.Open(global.SrcImagePath + global.SrcImageFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	dstfile, err := os.Create(
		global.DstImagePath + fmt.Sprintf(dstFileName,
			global.SrcImageFileName,
			"jpg",
		),
	)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	_, err = io.Copy(dstfile, file)
	return err
}
