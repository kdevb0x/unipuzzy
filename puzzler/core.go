package puzzler

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"os"

	"gocv.io/x/gocv"
)

// image of what the finished puzzle should look like (ie front of the box)
var referenceImage image.Image

type Grid struct {
	PieceCount     int // total n
	Xcount, Ycount int // number of rows (y) and columns (x)
	Pieces         []piece
}

type pos struct {
	xpos int
	ypos int
}

type piece struct {
	position  pos
	perimiter bool
	level     int // resolution level that piece
}
type View struct {
	ref         image.Image
	totalN      int
	overlayGrid Grid
}

func splitRefImage(imgpath string, numcols, numrow int) ([]piece, error) {
	var imgbounds = refimg.Bounds()
	var splitimg = make([]pieces, numcols*numrow)

	file, err := os.Open(imgpath)
	if err != nil {
		log.Printf("cat open ref image at %s, error: %s", imgpath, err)
		return nil, err
	}
	processed := image.Decode(file)

}
func prepareData(totalN int, finalState image.Image) (*View, error)
