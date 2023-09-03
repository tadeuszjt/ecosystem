package main

import (
	"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/tadeuszjt/geom/generic"
	"github.com/tadeuszjt/gfx"
)

var (
	circleTex  gfx.TexID
	botTex     gfx.TexID
	arrowTex   gfx.TexID
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
    frameRect.Min = geom.Vec2[float32]{}
	frameRect.Max = w.Size()
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
		mouseWorld = frameToWorld().TimesVec2(ev.Position, 1)

		if botHeld {
			for i, id := range bots.id {
				if id == botHeldId {
					bots.ori[i] = geom.Ori2[float32]{
						mouseWorld.X,
						mouseWorld.Y,
						bots.ori[i].Theta,
					}
					break
				}
			}
		} else if mouseHeld {
			camPos.PlusEquals(frameToWorld().TimesVec2(oldMousePos.Minus(mousePos), 0))
		}

	case gfx.MouseButton:
		if ev.Button == glfw.MouseButtonLeft {
			switch ev.Action {
			case glfw.Press:
				mouseHeld = true

				for i := range bots.ori {
					if bots.ori[i].Vec2().Minus(mouseWorld).Len() < 32 {
						fmt.Println("do_stuff")
						botHeld = true
						botHeldId = bots.id[i]
						break
					}
				}

			case glfw.Release:
				mouseHeld = false
				botHeld = false
			}
		}

	}
}

func keyboard(w *gfx.Win, ev gfx.KeyEvent) {
	switch ev.Key {
	case glfw.KeySpace:
		if ev.Action == glfw.Press {
			botsPause = !botsPause
		}
	}
}

func setup(w *gfx.Win) error {
	var err error

	circleTex, err = w.LoadTextureFromFile("circle.png")
	if err != nil {
		return err
	}

	botTex, err = w.LoadTextureFromFile("boat.png")
	if err != nil {
		return err
	}

	arrowTex, err = w.LoadTextureFromFile("arrow.png")
	if err != nil {
		return err
	}

	return err
}

func draw(w *gfx.Win, c gfx.Canvas) {
	arrows = nil
	update()
	mat := worldToFrame()
	drawBots(c, botTex, mat)
	drawArrows(c, arrowTex, mat)

	if botHeld {
		for i, id := range bots.id {
			if id == botHeldId {
				bots.brain[i].draw(c, circleTex)
				break
			}
		}
	}
}

func main() {
	start()

	gfx.RunWindow(gfx.WinConfig{
		SetupFunc:  setup,
		DrawFunc:   draw,
		ResizeFunc: size,
		MouseFunc:  mouse,
		KeyFunc:    keyboard,
		Title:      "Neurobots",
	})
}
