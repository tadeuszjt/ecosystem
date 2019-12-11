package main

import (
	"github.com/tadeuszjt/geom/32"
	"github.com/tadeuszjt/gfx"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	circleTex  gfx.TexID
	starTex    gfx.TexID
	frameRect  geom.Rect
	mousePos   geom.Vec2
	mouseWorld geom.Vec2
	mouseHeld  bool
	camPos     geom.Vec2
	camZoom    = float32(1)
)

func camRect() geom.Rect {
	return geom.RectCentredAt(
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
		oldMousePos := mousePos
		mousePos = ev.Position
		if mouseHeld {
			camPos.PlusEquals(frameToWorld().TimesVec2(oldMousePos.Minus(mousePos), 0).Vec2())
		}
		mouseWorld = frameToWorld().TimesVec2(ev.Position, 1).Vec2()
		
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
	
	circleTex, err = w.LoadTexture("circle.png")
	if err != nil {
		return err
	}
	
	starTex, err = w.LoadTexture("star.png")
	return err
}

func draw(w *gfx.WinDraw) {
	update()
	
	texCoords := [4]geom.Vec2{{0, 0}, {1, 0}, {1, 1}, {0, 1}}
	
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
			verts[j] = m.TimesVec2(verts[j], 1).Vec2()
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
	w.DrawVertexData(blobsData, &circleTex, &mat)
	w.DrawVertexData(predsData, &starTex, &mat)
	

	w.DrawText("Bingy", geom.Vec2{}, 80)
}

func main() {
	start()

	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		ResizeFunc: size,
		MouseFunc:  mouse,
		Title:      "Ecosystem",
	})
}
