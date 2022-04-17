package vec3

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (a *Vec3) Plus(v *Vec3) *Vec3 {
	a.X += v.X
	a.Y += v.Y
	a.Z += v.Z
	return a
}

func (a *Vec3) Inverse() *Vec3 {
	return a.Scale(-1)
}

func (a *Vec3) Scale(f float64) *Vec3 {
	a.X = f * a.X
	a.Y = f * a.Y
	a.Z = f * a.Z
	return a
}
