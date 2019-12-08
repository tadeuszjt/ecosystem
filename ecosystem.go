package main

import (
	"github.com/tadeuszjt/geom/geom32"
	"github.com/tadeuszjt/gfx"
	"github.com/tadeuszjt/data"
	"math/rand"
)

const (
	blobsStart       = 10
	blobsCollideDist = 20
	blobsBreedOdds   = 60
	blobsBreedDist   = 40
	blobsMaxAge      = 2000
	
	predsStart       = 3
	predsInitialFed  = 100
	predsEatRadius   = 30
	predsSightRadius = 100
	predsSpeed       = 0.6
)



var (
	arena = geom.RectCentred(1000, 1000)

	blobsT = data.Table{ &blobs.pos, &blobs.col, &blobs.age }
	blobs struct {
		pos geom.SliceVec2
		col SliceColour
		age data.SliceInt
	}
	
	predsT = data.Table{ &preds.ori, &preds.fed }
	preds struct {
		ori geom.SliceOri2
		fed data.SliceInt
	}
)

func randColour() gfx.Colour {
	return gfx.Colour{
		R: rand.Float32() * 0.5,
		G: rand.Float32() *0.4 + 0.6,
		B: rand.Float32() * 0.5,
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

func spawnBlob() {
	spawnPos := geom.Vec2Rand(arena)
	if !blobCollidesWithAny(spawnPos) {
		blobsT.Append(spawnPos, randColour(), 0)
	}
}

func spawnPred() {
	spawnPos := geom.Vec2Rand(arena)
	predsT.Append(
		geom.Ori2{spawnPos.X, spawnPos.Y, 0},
		predsInitialFed,
	)
}
	
func blobsInRange(pos geom.Vec2, radius float32) (s []geom.Vec2) {
	for _, p := range blobs.pos {
		if p.Minus(pos).Len2() < radius*radius {
			s = append(s, p)
		}
	}
	
	return s
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
	/* breed blobs */
	for i, pos := range blobs.pos {
		if rand.Intn(blobsBreedOdds) == 0 {
			dist := rand.Float32() * blobsBreedDist
			childPos := pos.Plus(geom.Vec2RandNormal().ScaledBy(dist))
			
			if arena.Contains(childPos) && !blobCollidesWithAny(childPos) {
				blobsT.Append(childPos, blobs.col[i], 0)
			}
		}
	}
	
	/* blobs die */
	blobsT.Filter(func(i int) bool {
		return rand.Intn(blobsMaxAge - blobs.age[i]) != 0
	})
	
	/* predators move */
	for i := range preds.ori {
		pos := preds.ori[i].Vec2()
		blobsSee := blobsInRange(pos, predsSightRadius)
		
		var vel geom.Vec2
		for _, bpos := range blobsSee {
			toBlob := bpos.Minus(pos)
			distToBlob := toBlob.Len()
			
			add := toBlob.Normal().ScaledBy(predsSightRadius - distToBlob)
			vel.PlusEquals(add)
		}
		
		preds.ori[i].PlusEquals(
			vel.Normal().ScaledBy(predsSpeed).Ori2(),
		)
	}
	
	/* predators eat */
	for i := range preds.ori {
		pos := preds.ori[i].Vec2()
		
		blobsT.Filter(func(j int) bool {
			if blobs.pos[j].Minus(pos).Len2() < predsEatRadius*predsEatRadius {
				preds.fed[i]++
				if preds.fed[i] > 100 {
					preds.fed[i] = 100
				}
				
				return false
			}
			
			return true
		})
	}
}
