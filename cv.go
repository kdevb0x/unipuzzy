package puzzler

import (
	"io"
	"os"
	"image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

// image of what the finished puzzle should look like (ie front of the box)
var referenceImage image.Image

type Grid
type View struct {
	ref image.Image
	totalN int
	overlayGrid Grid
}
func prepareData(totalN int, finalState image.Image) (*View, error)
