package puzzler

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

/*

To make it easier to reason about, we take the finished puzzle,
(thats all n pieces, in their final destination)
and split it up into a grid of n positions.

We label the X for the horizontal, and y for the vertical.
So the left-most, postition on the bottom-most row is at {X0, Y0}

Rough Visualization:

   /----------------------------------\
y27|				      |
y26|				      |
.  |				      |
.  |                                  |
.  |  				      |
.  |  // Amazing Puzzle Scene Here    |
.  | 				      |
.  |                                  |
y3 |				      |
y2 |				      |
y1 |				      |
y0 \__________________________________/
  x0x1x2x3x4x5x6x7x8x9x10x11x12x13....x37


*/

const (
	// X0 is leftmost size of grid, X27 is rightmost
	parimeter piece = iota
	X0
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
	X16
	X17
	X18
	X19
	X20
	X21
	X22
	X23
	X24
	X25
	X26
	X27

	// Y0 is bottommost row, Y27 is the top
	Y0
	Y1
	Y2
	Y3
	Y4
	Y5
	Y6
	Y7
	Y8
	Y9
	Y10
	Y11
	Y12
	Y13
	Y14
	Y15
	Y16
	Y17
	Y18
	Y19
	Y20
	Y21
	Y22
	Y23
	Y24
	Y25
	Y26
	Y27
)

type pieLce struct {
	estimatedFinPos area
}

func bruteforce(timer time.Timer, pieces int) time.Duration {
	var n int = pieces
	var countLside, countRside, countTop, countBottom int
	countLside, countRside = 27, 27
	countTop, countBottom = 36, 36
	var corners = 4
	var finalPos []pieces
	for i := 0; i < n; i++ {

	}
}
