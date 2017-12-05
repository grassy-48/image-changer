package color

import "image"

type Converter interface {
	Convert(img image.Image) error
}
