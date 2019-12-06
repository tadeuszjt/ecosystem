package main

import (
	"github.com/tadeuszjt/geom/geom32"
	"github.com/tadeuszjt/gfx"
)

const startBlobs = 10

var (
	arena = geom.RectCentred(1000, 1000)

	blobs, preds struct{
		pos []geom.Vec2
		col []gfx.Colour
		age []int
	}
)


func start() {
	for i := 0; i < startBlobs; i++ {
	}
}
