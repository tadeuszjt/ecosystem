package main

import (
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
	"math/rand"
)

const (
	blobsStart       = 200
	blobsCollideDist = 20
	blobsBreedOdds   = 60
	blobsBreedDist   = 40
	blobsMaxAge      = 2000

	predsStart       = 5
	predsStartFed    = 100.
	predsChildFed    = 20.
	predsMaxFed      = 100.
	predsEatPlus     = 1.6
	predsFedBleed    = 0.3
	predsEatRadius   = 30
	predsSightRadius = 200
	predsSpeed       = 2.6
	predsBreedOdds   = 1400
)

var (
	arena = geom.RectCentred(2000, 2000)

	blobsT = data.Table{&blobs.pos, &blobs.col, &blobs.age}
	blobs  struct {
		pos geom.SliceVec2
		col data.SliceT[gfx.Colour]
		age data.SliceT[int]
	}

	predsT = data.Table{&preds.ori, &preds.dir, &preds.fed}
	preds  struct {
		ori geom.SliceOri2
		dir geom.SliceVec2
		fed data.SliceT[float32]
	}
)

func randColour() gfx.Colour {
	return gfx.Colour{
		R: rand.Float32() * 0.5,
		G: rand.Float32()*0.4 + 0.6,
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
	spawnPos := geom.Vec2Rand(arena).Ori2()
	spawnDir := geom.Vec2RandNormal()
	predsT.Append(spawnPos, spawnDir, float32(predsStartFed))
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

	/* breed preds */
	for i := range preds.ori {
		if rand.Intn(predsBreedOdds) == 0 {
			predsT.Append(
				preds.ori[i].Vec2().Plus(geom.Vec2RandNormal().ScaledBy(10)).Ori2(),
				geom.Vec2RandNormal(),
				float32(predsChildFed),
			)
		}
	}

	/* blobs die */
	blobsT.Filter(func(i int) bool {
		return rand.Intn(blobsMaxAge-blobs.age[i]) != 0
	})

	/* predators die */
	predsT.Filter(func(i int) bool {
		preds.fed[i] -= predsFedBleed
		return preds.fed[i] > 0.0
	})

	/* predators move */
	for i := range preds.ori {
		pos := preds.ori[i].Vec2()
		vel := preds.dir[i]
		blobsSee := blobsInRange(pos, predsSightRadius)

		for _, bpos := range blobsSee {
			toBlob := bpos.Minus(pos)
			distToBlob := toBlob.Len()
			scalar := predsSightRadius - distToBlob
			scalar *= scalar

			add := toBlob.Normal().ScaledBy(scalar)
			vel.PlusEquals(add)
		}

		preds.dir[i] = preds.dir[i].Plus(vel.Normal().ScaledBy(0.1)).Normal()
		newPos := pos.Plus(preds.dir[i].ScaledBy(predsSpeed))

		if newPos.X > arena.Max.X {
			newPos.X = arena.Max.X
		} else if newPos.X < arena.Min.X {
			newPos.X = arena.Min.X
		}

		if newPos.Y > arena.Max.Y {
			newPos.Y = arena.Max.Y
		} else if newPos.Y < arena.Min.Y {
			newPos.Y = arena.Min.Y
		}

		preds.ori[i] = newPos.Ori2()
	}

	/* predators eat */
	for i := range preds.ori {
		pos := preds.ori[i].Vec2()

		blobsT.Filter(func(j int) bool {
			if blobs.pos[j].Minus(pos).Len2() < predsEatRadius*predsEatRadius {

				preds.fed[i] += predsEatPlus
				if preds.fed[i] > predsMaxFed {
					preds.fed[i] = predsMaxFed
				}

				return false
			}

			return true
		})
	}
}
