package main

import (
	//"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tadeuszjt/data"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

var (
	frameRect geom.Rect[float32]
	camPos    geom.Vec2[float32]
	camZoom   = float32(0.01)

	mousePos       geom.Vec2[float32]
	mouseWorld     geom.Vec2[float32]
	mouseHeld      bool
	bodyHeld       = data.KeyInvalid
	bodyHeldOffset geom.Vec2[float32]

	timeStep = 1. / 60.
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
				for i := range bodies.physKeys {
					ori := geom.Ori2Convert[float64, float32](world.GetOrientation(bodies.physKeys[i]))
					trans := geom.Mat3Translation(ori.Vec2().ScaledBy(-1))
					rot := geom.Mat3Rotation(-ori.Theta)
					mouseBody := rot.Product(trans).TimesVec2(mouseWorld, 1)

					if poly.Contains(mouseBody) {
						bodyHeld = bodies.physKeys[i]
						bodyHeldOffset = mouseBody
						mouseHeld = false
					}
				}

			case glfw.Release:
				bodyHeld = data.KeyInvalid
				mouseHeld = false
			}
		}

	}
}

func setup(w *gfx.Win) error {
	addBody(geom.Ori2[float64]{1, 1, 1})
	addBody(geom.Ori2[float64]{2, 2, 2})
	addFish(geom.Vec2[float64]{2, 4})
	addFish(geom.Vec2[float64]{4, 1})

	return nil
}

func draw(w *gfx.Win, c gfx.Canvas) {
	if bodyHeld != data.KeyInvalid {
		ori := world.GetOrientation(bodyHeld)
		offset := geom.Vec2Convert[float32, float64](bodyHeldOffset).RotatedBy(ori.Theta)
		delta := geom.Vec2Convert[float32, float64](mouseWorld).Minus(ori.Vec2().Plus(offset))
		world.ApplyImpulse(bodyHeld, delta, offset, timeStep)
	}

	world.Update(timeStep)

	mat := worldToFrame()
	gfx.DrawSprite(c, geom.Ori2[float32]{}, geom.RectOrigin[float32](1, 1), gfx.Red, &mat, nil)
	drawBodies(c, mat)
}

func main() {
	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		ResizeFunc: size,
		MouseFunc:  mouse,
		Width:      1024,
		Height:     800,
		Title:      "fish",
	})
}
