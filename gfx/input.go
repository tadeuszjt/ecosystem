package gfx

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/tadeuszjt/ecosystem/geom/32"
)

type MouseEvent interface {
}

type MouseScroll struct {
	Dx, Dy float32
}

type MouseMove struct {
	Position geom.Vec2
}

type MouseButton struct {
	Button glfw.MouseButton
	Action glfw.Action
	Mods   glfw.ModifierKey
}
