package main

import (
	"github.com/tadeuszjt/geom/geom32"
	"github.com/tadeuszjt/gfx"
)

var (
	texID        gfx.TexID
	frameRect    geom.Rect
	mousePos     geom.Vec2
	mouseWorld   geom.Vec2
	camPos       geom.Vec2
	camZoom      = float32(1)
	
	worldToFrame = geom.Mat3Identity()
	frameToWorld = geom.Mat3Identity()
)

func updateMatrices() {
	camRect := geom.RectCreate(
		frameRect.Width() * camZoom,
		frameRect.Height() * camZoom,
		camPos)
	worldToFrame = geom.Mat3Camera2D(camRect, frameRect)
	frameToWorld = geom.Mat3Camera2D(frameRect, camRect)
}

func size(w, h int) {
	frameRect = geom.RectOrigin(float32(w), float32(h))
	updateMatrices()
}

func mouse(w *gfx.Win, event gfx.MouseEvent) {
	switch ev := event.(type) {
	case gfx.MouseScroll:
		oldPos := mouseWorld
		camZoom *= 1 - 0.04*ev.Dy
		updateMatrices()
		newPos := frameToWorld.TimesVec2(mousePos, 1).Vec2()
		camPos.PlusEquals(oldPos.Minus(newPos))
		updateMatrices()
	
	case gfx.MouseMove:
		mousePos = ev.Position
		mouseWorld = frameToWorld.TimesVec2(ev.Position, 1).Vec2()
	}	
}

func setup(w *gfx.Win) error {
	var err error
	texID, err = w.LoadTexture("circle.png")
	return err
}

func draw(w *gfx.WinDraw) {
	w.DrawRect(
		geom.RectCreate(300, 300, mouseWorld),
		&texID,
		&gfx.Colour{0, 1, 0, 1},
		&worldToFrame,
	)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc: setup,
		DrawFunc:  draw,
		ResizeFunc: size,
		MouseFunc: mouse,
		Title:     "Benis",
	})
}
