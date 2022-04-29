package vec3

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// type Point3  Vec3
// type Color  Vec3

func (a *Vec3) Plus(v *Vec3) *Vec3 {
	a.X += v.X
	a.Y += v.Y
	a.Z += v.Z
	return a
}
func (a *Vec3) Minus(v *Vec3) *Vec3 {
	a.X -= v.X
	a.Y -= v.Y
	a.Z -= v.Z
	return a
}

func (a *Vec3) Times(v *Vec3) *Vec3 {
	a.X *= v.X
	a.Y *= v.Y
	a.Z *= v.Z
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

func (a *Vec3) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a *Vec3) Length() float64 {
	return math.Sqrt(a.LengthSquared())
}

func (a *Vec3) Dot() float64 {
	return a.LengthSquared()
}

func (a *Vec3) Cross(t *Vec3) *Vec3 {
	x := a.Y*t.Z - a.Z*t.Y
	y := a.Z*t.X - a.X*t.Z
	z := a.X*t.Y - a.Y*t.X
	a.X = x
	a.Y = y
	a.Z = z
	return a
}

func (a *Vec3) UnitVector() *Vec3 {
	return a.Scale(1 / a.Length())
}
