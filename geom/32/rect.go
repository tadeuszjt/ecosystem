package geom

type Rect struct {
	Min, Max Vec2
}

func RectOrigin(w, h float32) Rect {
	return Rect{
		Vec2{0, 0},
		Vec2{w, h},
	}
}

func RectCentred(w, h float32) Rect {
	wh := w / 2
	hh := h / 2
	return Rect{
		Vec2{-wh, -hh},
		Vec2{wh, hh},
	}
}

func RectCentredAt(w, h float32, pos Vec2) Rect {
	wh := w / 2
	hh := h / 2
	return Rect{
		Vec2{pos.X - wh, pos.Y - hh},
		Vec2{pos.X + wh, pos.Y + hh},
	}
}

func MakeRect(w, h float32, pos Vec2) Rect {
	return Rect{
		Vec2{pos.X, pos.Y},
		Vec2{pos.X + w, pos.Y + h},
	}
}

func (r Rect) Width() float32 {
	return r.Max.X - r.Min.X
}

func (r Rect) Height() float32 {
	return r.Max.Y - r.Min.Y
}

func (r Rect) Contains(v Vec2) bool {
	return v.X >= r.Min.X &&
		v.X <= r.Max.X &&
		v.Y >= r.Min.Y &&
		v.Y <= r.Max.Y
}

func (r Rect) Verts() [4]Vec2 {
	return [4]Vec2{
		r.Min,
		{r.Max.X, r.Min.Y},
		r.Max,
		{r.Min.X, r.Max.Y},
	}
}
