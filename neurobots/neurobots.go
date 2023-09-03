package main

import (
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
	"github.com/tadeuszjt/neuralnetwork"
    "math/rand"
    "math"
)

const (
	botsStart       = 100
	botsSightRadius = 200
	botsSightWidth  = 2
	botsSpeed       = 2
	botsBreedOdds   = 1400
	botsNumSensors  = 8
	botsSize        = 20
)

var (
	arena = geom.RectCentred[float32](2000, 2000)

    bots struct {
        data.Table
        ori   data.RowT[geom.Ori2[float32]]
        dir   data.RowT[geom.Vec2[float32]]
        col   data.RowT[gfx.Colour]
        id    data.RowT[int]
        brain data.RowT[BotBrain]
    }

	botIds    = 0
	botHeld   = false
	botHeldId = 0
	botsPause = true
)

func init() {
    bots.Table = data.Table{
        &bots.ori,
        &bots.dir,
        &bots.col,
        &bots.id,
        &bots.brain,
    }
}

func addBot(ori geom.Ori2[float32], dir geom.Vec2[float32], col gfx.Colour) {
	nn := nn.MakeNeuralNetwork(botsNumSensors, 2, 3, 20)
	nn.RandomiseWeights()

	bots.Append(ori, dir, col, botIds, BotBrain{
		network: nn,
	})
	botIds++
}

func spawnBot() {
	spawnPos := geom.Vec2Rand(arena)
	spawnOri := geom.Ori2[float32]{spawnPos.X, spawnPos.Y, rand.Float32() * 2 * math.Pi}
	spawnDir := geom.Vec2RandNormal[float32]()
	addBot(spawnOri, spawnDir, gfx.Colour{1, 1, 1, 1})
}

func start() {
	for i := 0; i < botsStart; i++ {
		spawnBot()
	}
}

func update() {
	// calculate nn inputs from sensors
	for i := range bots.ori {
		bots.brain[i].network.ClearInputs()

		for j := range bots.ori {
			if i == j {
				continue
			}

			delta := bots.ori[j].Vec2().Minus(bots.ori[i].Vec2())

			if delta.Len2() > botsSightRadius*botsSightRadius {
				continue
			}

			bearing := delta.Theta() - bots.ori[i].Theta

			if bearing > botsSightWidth || -bearing > botsSightWidth {
				continue
			}

			sensorIdx := int(botsNumSensors/2*bearing/botsSightWidth) + botsNumSensors/2
			activation := (botsSightRadius - delta.Len()) / botsSightRadius

			inputs := bots.brain[i].network.Inputs()
			if activation > bots.brain[i].sensors[sensorIdx] {
				inputs[sensorIdx] = activation
			}
		}
	}

	// control bots from nn outputs
	for i := range bots.dir {
		bots.brain[i].network.Process()
		output := bots.brain[i].network.Outputs()[1] - bots.brain[i].network.Outputs()[0]
		bots.dir[i] = bots.dir[i].RotatedBy(output * 0.01)
		bots.ori[i].Theta = bots.dir[i].Theta()
	}

    // kill bots
	for i := 0; i < len(bots.ori); i++ {
		for j := 0; j < len(bots.ori); j++ {
			if i == j {
				continue
			}

			delta := bots.ori[i].Vec2().Minus(bots.ori[j].Vec2())
			if delta.Len2() < botsSize*botsSize {
				bots.Delete(i)
				bots.Delete(j)
				break
			}
		}
	}

	if !botsPause {
		for i := range bots.ori {
			bots.ori[i].PlusEquals(bots.dir[i].ScaledBy(botsSpeed).Ori2())
		}

        // breed bots
        for i := 0; i < len(bots.ori); i++ {
            if rand.Intn(botsBreedOdds) == 0 {
                dir := geom.Vec2RandNormal[float32]()
                ori := geom.Ori2[float32]{
                    bots.ori[i].Vec2().Plus(dir.ScaledBy(botsSize)).X,
                    bots.ori[i].Vec2().Plus(dir.ScaledBy(botsSize)).Y,
                    dir.Theta(),
                }
                addBot(ori, dir, bots.col[i])
            }
        }
    }

	// clamp to arena
	for i := range bots.ori {
		if bots.ori[i].X > arena.Max.X {
			bots.ori[i].X = arena.Max.X
			bots.dir[i].X *= -1
		} else if bots.ori[i].X < arena.Min.X {
			bots.ori[i].X = arena.Min.X
			bots.dir[i].X *= -1
		}
		if bots.ori[i].Y > arena.Max.Y {
			bots.ori[i].Y = arena.Max.Y
			bots.dir[i].Y *= -1
		} else if bots.ori[i].Y < arena.Min.Y {
			bots.ori[i].Y = arena.Min.Y
			bots.dir[i].Y *= -1
		}
	}
}

func drawBots(c gfx.Canvas, tex gfx.TexID, mat geom.Mat3[float32]) {
	texCoords := [4]geom.Vec2[float32]{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	botsData := make([]float32, 0, 6*8*len(bots.ori))

	for i, ori := range bots.ori {
		botsCol := bots.col[i]
		verts := geom.RectCentred[float32](botsSize, botsSize).Verts()

		for j := range verts {
			verts[j] = ori.Mat3Transform().TimesVec2(verts[j], 1)
		}

		for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
			botsData = append(
				botsData,
				verts[j].X, verts[j].Y,
				texCoords[j].X, texCoords[j].Y,
				botsCol.R, botsCol.G, botsCol.B, botsCol.A,
			)
		}
	}

	c.Draw2DVertexData(botsData, &tex, &mat)
}
