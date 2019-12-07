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
	
	predsStart       = 3
	predsInitialFed  = 100
	predsEatRadius   = 30
)

var (
	arena = geom.RectCentred(1000, 1000)

	blobs struct {
		pos []geom.Vec2
		col []gfx.Colour
		age []int
	}
	
	preds struct {
		ori []geom.Ori2
		fed []int
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
	if !blobCollidesWithAny(spawnPos) {
		addBlob(spawnPos, randColour())
	}
}

func killBlob(i int) {
	end := len(blobs.pos) - 1
	
	if i < end {
		blobs.pos[i] = blobs.pos[end]
		blobs.col[i] = blobs.col[end]
		blobs.age[i] = blobs.age[end]
	}
	
	blobs.pos = blobs.pos[:end]
	blobs.col = blobs.col[:end]
	blobs.age = blobs.age[:end]
}


func addPred(pos geom.Vec2) {
	preds.ori = append(preds.ori, geom.Ori2{pos.X, pos.Y, 0})
	preds.fed = append(preds.fed, predsInitialFed)
}

func spawnPred() {
	addPred(geom.Vec2Rand(arena))
}
	

func start() {
	for i := 0; i < blobsStart; i++ {
		spawnBlob()
	}
	
	for i := 0; i < predsStart; i++ {
		spawnPred()
	}
}

func update() {
	if rand.Intn(blobsSpawnOdds) == 0 {
		spawnBlob()
	}
	
	/* predators move */
	for i := range preds.ori {
	}
	
	/* predators eat */
	for i := range preds.ori {
		pos := preds.ori[i].Vec2()
		
		for j := 0; j < len(blobs.pos); j++ {
			if blobs.pos[j].Minus(pos).Len2() < predsEatRadius*predsEatRadius {
				killBlob(j)
				j--
				preds.fed[i]++
				
				if preds.fed[i] > 100 {
					preds.fed[i] = 100
				}
			}
		}
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
