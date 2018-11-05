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

func gridSize(nsize int, xcount, ycount ...int) *Grid {
	var grid = &Grid{
		PieceCount: nsize,
		Xcount:     xcount,
		Ycount:     ycount,
	}

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
