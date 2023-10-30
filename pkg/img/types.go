package img

import "image"

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

type RGB struct {
	R float64
	G float64
	B float64
}
