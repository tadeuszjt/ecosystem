package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

var (
	circleTex gfx.TexID
	starTex   gfx.TexID
	//text       gfx.Text
	frameRect  geom.Rect[float32]
	mousePos   geom.Vec2[float32]
	mouseWorld geom.Vec2[float32]
	mouseHeld  bool
	camPos     geom.Vec2[float32]
	camZoom    = float32(1)
)

func camRect() geom.Rect[float32] {
	return geom.RectCentredAt(
		frameRect.Width()*camZoom,
		frameRect.Height()*camZoom,
		camPos)
}

func worldToFrame() geom.Mat3[float32] {
	return geom.Mat3Camera2D(camRect(), frameRect)
}

func frameToWorld() geom.Mat3[float32] {
	return geom.Mat3Camera2D(frameRect, camRect())
}

func size(w *gfx.Win) {
	v := w.Size()
	frameRect = geom.RectOrigin(float32(v.X), float32(v.Y))
}

func mouse(w *gfx.Win, event gfx.MouseEvent) {
	switch ev := event.(type) {
	case gfx.MouseScroll:
		oldPos := mouseWorld
		camZoom *= 1 - 0.04*ev.Dy
		newPos := frameToWorld().TimesVec2(mousePos, 1)
		camPos.PlusEquals(oldPos.Minus(newPos))

	case gfx.MouseMove:
		oldMousePos := mousePos
		mousePos = ev.Position
		if mouseHeld {
			camPos.PlusEquals(frameToWorld().TimesVec2(oldMousePos.Minus(mousePos), 0))
		}
		mouseWorld = frameToWorld().TimesVec2(ev.Position, 1)

	case gfx.MouseButton:
		if ev.Button == glfw.MouseButtonLeft {
			switch ev.Action {
			case glfw.Press:
				mouseHeld = true
			case glfw.Release:
				mouseHeld = false
			}
		}

	}
}

func setup(w *gfx.Win) error {
	var err error

	//	text.SetString("good morning sir!")
	//	text.SetSize(12)

	circleTex, err = w.LoadTextureFromFile("circle.png")
	if err != nil {
		return err
	}

	starTex, err = w.LoadTextureFromFile("star.png")
	return err
}

func draw(w *gfx.Win, c gfx.Canvas) {
	update()
	update()
	update()
	update()
	update()
	update()
	update()

	texCoords := [4]geom.Vec2[float32]{{0, 0}, {1, 0}, {1, 1}, {0, 1}}

	blobsSize := float32(20)
	blobsData := make([]float32, 0, 6*8*len(blobs.pos))

	for i, pos := range blobs.pos {
		verts := geom.RectCentredAt(blobsSize, blobsSize, pos).Verts()

		for _, j := range []int{0, 1, 2, 0, 2, 3} {
			col := blobs.col[i]
			blobsData = append(blobsData,
				verts[j].X, verts[j].Y,
				texCoords[j].X, texCoords[j].Y,
				col.R, col.G, col.B, col.A,
			)
		}
	}

	predsSize := float32(40)
	predsData := make([]float32, 0, 6*8*len(preds.ori))
	predsCol := gfx.Colour{1, 1, 1, 1}

	for _, ori := range preds.ori {
		m := ori.Mat3Transform()

		verts := geom.RectCentred(predsSize, predsSize).Verts()

		for j := range verts {
			verts[j] = m.TimesVec2(verts[j], 1)
		}

		for _, j := range [6]int{0, 1, 2, 0, 2, 3} {
			predsData = append(predsData,
				verts[j].X, verts[j].Y,
				texCoords[j].X, texCoords[j].Y,
				predsCol.R, predsCol.G, predsCol.B, predsCol.A,
			)
		}
	}

	mat := worldToFrame()
	c.Draw2DVertexData(blobsData, &circleTex, &mat)
	c.Draw2DVertexData(predsData, &starTex, &mat)

	//	for i, ori := range preds.ori {
	//		text.SetString(fmt.Sprintf("%f", preds.fed[i]))
	//		w.DrawText(&text,
	//			worldToFrame().TimesVec2(ori.Vec2().Plus(
	//				geom.Vec2{0, predsSize / 2}), 1).Vec2())
	//	}
}

func main() {
	start()

	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		ResizeFunc: size,
		MouseFunc:  mouse,
		Width:      1024,
		Height:     800,
		Title:      "Ecosystem",
	})
}
