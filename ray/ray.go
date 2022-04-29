package ray

import "goRayTrace/vec3"

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
}

func (a Ray) At(t float64) *vec3.Vec3 {
	// a.Origin + t*a.Direction
	return a.Origin.Plus(a.Direction.Scale(t))
}
