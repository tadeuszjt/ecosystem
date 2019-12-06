package main

import (
	"github.com/tadeuszjt/geom/geom32"
	"github.com/tadeuszjt/gfx"
	"math/rand"
)

const (
	blobsStart       = 10
	blobsSpawnOdds   = 100
	blobsCollideDist = 20
	blobsBreedOdds   = 100
	blobsBreedDist   = 40
)

var (
	arena = geom.RectCentred(1000, 1000)

	blobs, preds struct {
		pos []geom.Vec2
		col []gfx.Colour
		age []int
	}
)

func randColour() gfx.Colour {
	return gfx.Colour{
		R: rand.Float32(),
		G: rand.Float32(),
		B: rand.Float32(),
		A: 1,
	}
}

func blobsCollide(a, b geom.Vec2) bool {
	sq := blobsCollideDist * blobsCollideDist
	return a.Minus(b).Len2() < float32(sq)
}

func blobCollidesWithAny(pos geom.Vec2) bool {
	for _, p := range blobs.pos {
		if blobsCollide(p, pos) {
			return true
		}
	}
	
	return false
}

func addBlob(pos geom.Vec2, col gfx.Colour) {
	blobs.pos = append(blobs.pos, pos)
	blobs.col = append(blobs.col, col)
	blobs.age = append(blobs.age, 0)
}

func spawnBlob() {
	spawnPos := geom.Vec2Rand(arena)
	if blobCollidesWithAny(spawnPos) {
		return
	}
	
	addBlob(spawnPos, randColour())
}

func start() {
	for i := 0; i < blobsStart; i++ {
		spawnBlob()
	}
}

func update() {
	if rand.Intn(blobsSpawnOdds) == 0 {
		spawnBlob()
	}
	
	/* breed blobs */
	for i, pos := range blobs.pos {
		if rand.Intn(blobsBreedOdds) == 0 {
			dist := rand.Float32() * blobsBreedDist
			childPos := pos.Plus(geom.Vec2RandNormal().ScaledBy(dist))
			
			if arena.Contains(childPos) && !blobCollidesWithAny(childPos) {
				addBlob(childPos, blobs.col[i])
			}
		}
	}
	
	
}
