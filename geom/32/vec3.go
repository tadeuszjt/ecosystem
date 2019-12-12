package geom

type Vec3 struct {
	X, Y, Z float32
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v.X, v.Y}
}

func (a Vec3) Dot(b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vec3) Times(b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (v Vec3) ScaledBy(f float32) Vec3 {
	return Vec3{f * v.X, f * v.Y, f * v.Z}
}

func (v Vec3) Ori2() Ori2 {
	return Ori2{v.X, v.Y, v.Z}
}
