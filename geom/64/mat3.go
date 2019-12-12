package geom

type Mat3 [9]float64

func Mat3Identity() Mat3 {
	return Mat3{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	}
}

func (a Mat3) Product(b Mat3) Mat3 {
	return Mat3{
		a[0]*b[0] + a[1]*b[3] + a[2]*b[6],
		a[0]*b[1] + a[1]*b[4] + a[2]*b[7],
		a[0]*b[2] + a[1]*b[5] + a[2]*b[8],

		a[3]*b[0] + a[4]*b[3] + a[5]*b[6],
		a[3]*b[1] + a[4]*b[4] + a[5]*b[7],
		a[3]*b[2] + a[4]*b[5] + a[5]*b[8],

		a[6]*b[0] + a[7]*b[3] + a[8]*b[6],
		a[6]*b[1] + a[7]*b[4] + a[8]*b[7],
		a[6]*b[2] + a[7]*b[5] + a[8]*b[8],
	}
}

func (m Mat3) TimesVec2(v Vec2, bias float64) Vec3 {
	return Vec3{
		X: m[0]*v.X + m[1]*v.Y + m[2]*bias,
		Y: m[3]*v.X + m[4]*v.Y + m[5]*bias,
		Z: m[6]*v.X + m[7]*v.Y + m[8]*bias,
	}
}

func Mat3Camera2D(camera, display Rect) Mat3 {
	sx := display.Width() / camera.Width()
	sy := display.Height() / camera.Height()

	tx := display.Min.X - sx*camera.Min.X
	ty := display.Min.Y - sy*camera.Min.Y

	return Mat3{
		sx, 0, tx,
		0, sy, ty,
		0, 0, 1,
	}
}
