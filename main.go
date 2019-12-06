package main

import (
	"github.com/tadeuszjt/geom/geom32"
	"github.com/tadeuszjt/gfx"
)

var (
	texID      gfx.TexID
	frameRect  geom.Rect
	mousePos   geom.Vec2
	mouseWorld geom.Vec2
	camPos     geom.Vec2
	camZoom    = float32(1)
)

func camRect() geom.Rect {
	return geom.RectCreate(
		frameRect.Width()*camZoom,
		frameRect.Height()*camZoom,
		camPos)
}

func worldToFrame() geom.Mat3 {
	return geom.Mat3Camera2D(camRect(), frameRect)
}

func frameToWorld() geom.Mat3 {
	return geom.Mat3Camera2D(frameRect, camRect())
}

func size(w, h int) {
	frameRect = geom.RectOrigin(float32(w), float32(h))
}

func mouse(w *gfx.Win, event gfx.MouseEvent) {
	switch ev := event.(type) {
	case gfx.MouseScroll:
		oldPos := mouseWorld
		camZoom *= 1 - 0.04*ev.Dy
		newPos := frameToWorld().TimesVec2(mousePos, 1).Vec2()
		camPos.PlusEquals(oldPos.Minus(newPos))

	case gfx.MouseMove:
		mousePos = ev.Position
		mouseWorld = frameToWorld().TimesVec2(ev.Position, 1).Vec2()
	}
}

func setup(w *gfx.Win) error {
	var err error
	texID, err = w.LoadTexture("circle.png")
	return err
}

func draw(w *gfx.WinDraw) {
	update()
	
	blobsSize := float32(10)
	blobsData := make([]float32, 0, 6*8*len(blobs.pos))
	texCoords := [4]geom.Vec2{{0, 0}, {1, 0}, {1, 1}, {0, 1}}

	for i, pos := range blobs.pos {
		verts := [4]geom.Vec2{
			pos.Plus(geom.Vec2{-blobsSize, -blobsSize}),
			pos.Plus(geom.Vec2{blobsSize, -blobsSize}),
			pos.Plus(geom.Vec2{blobsSize, blobsSize}),
			pos.Plus(geom.Vec2{-blobsSize, blobsSize}),
		}

		for _, j := range []int{0, 1, 2, 0, 2, 3} {
			col := blobs.col[i]
			blobsData = append(blobsData,
				verts[j].X, verts[j].Y,
				texCoords[j].X, texCoords[j].Y,
				col.R, col.G, col.B, col.A,
			)
		}
	}

	mat := worldToFrame()
	w.DrawVertexData(blobsData, &texID, &mat)
}

func main() {
	start()

	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		ResizeFunc: size,
		MouseFunc:  mouse,
		Title:      "Benis",
	})
}
