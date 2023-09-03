package main

import (
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
	"github.com/tadeuszjt/neuralnetwork"
)

type BotBrain struct {
	sensors [botsNumSensors]float32
	network nn.NeuralNetwork
}

type SliceBrain []BotBrain

func (s *SliceBrain) Len() int {
	return len(*s)
}

func (s *SliceBrain) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *SliceBrain) Delete(i int) {
	end := s.Len() - 1
	if i < end {
		s.Swap(i, end)
	}

	*s = (*s)[:end]
}

func (s *SliceBrain) Append(t interface{}) {
	i, ok := t.(BotBrain)
	if !ok {
		panic("wrong type")
	}

	*s = append(*s, i)
}

func (b *BotBrain) draw(c gfx.Canvas, circleTex gfx.TexID) {
	spriteRect := geom.RectOrigin[float32](20, 20)

	for i, input := range b.network.Inputs() {
		ori := geom.Ori2[float32]{0, float32(i) * 30, 0}
		col := gfx.Colour{input, 0, 0, 1}
		gfx.DrawSprite(c, ori, spriteRect, col, nil, &circleTex)
	}

	for l := 0; l < b.network.NumLayers; l++ {
		for n := 0; n < b.network.NumNeuronsPerLayer; n++ {
			ori := geom.Ori2[float32]{float32((l + 1) * 30), float32(n * 30), 0}
			col := gfx.Colour{b.network.Neurons()[b.network.NumNeuronsPerLayer*l+n], 0, 0, 1}
			gfx.DrawSprite(c, ori, spriteRect, col, nil, &circleTex)
		}
	}

	for i, output := range b.network.Outputs() {
		ori := geom.Ori2[float32]{float32((b.network.NumLayers + 1) * 30), float32(i) * 30, 0}
		col := gfx.Colour{output, 0, 0, 1}
		gfx.DrawSprite(c, ori, spriteRect, col, nil, &circleTex)
	}
}
