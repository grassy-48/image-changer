package main

import (
	"image-changer/actions/color"
	"image-changer/actions/extention"
	"image-changer/actions/origin"

	"os"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	version = "0.0.1"
	appName = "image-changer"
)

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = version
	app.Compiled = time.Now()

	app.Commands = []cli.Command{
		{
			Name:   "origin",
			Usage:  "そのままの画像を出力する",
			Action: origin.Origin,
		},
		{
			Name:   "extention",
			Usage:  "画像の拡張子を変えて出力する",
			Action: extention.Extention,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "ext, e",
					Usage: "出力する拡張子(jpg(jpeg)), gif, pngのみ)",
				},
			},
		},
		{
			Name:   "color",
			Usage:  "画像の色を変更する",
			Action: color.Color,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "type, t",
					Usage: "変更する色のタイプ",
				},
			},
		},
	}
	app.Run(os.Args)
}
