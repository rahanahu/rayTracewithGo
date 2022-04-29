package main

import (
	"fmt"
	"goRayTrace/ray"
	"goRayTrace/vec3"
	"os"
)

// image
const aspect_ratio float64 = 16.0 / 9.0
const image_width int = 400
const image_height int = int(float64(image_width) / aspect_ratio)

// Camera
var viewport_hight = 2.0
var viewport_witdh = 400
var focal_length = 1.0
var origin = vec3.Vec3{X: 0, Y: 0, Z: 0}
var horizontal = vec3.Vec3{X: float64(viewport_witdh), Y: 0, Z: 0}
var vertical = vec3.Vec3{X: 0, Y: float64(viewport_hight), Z: 0}

// lower_left_corner = origin - horizontal/2 - vertical/2 - vec3(0, 0, focal_length)
var lower_left_corner = origin.Minus(horizontal.Scale(0.5)).Minus(vertical.Scale(0.5)).Minus(&vec3.Vec3{X: 0, Y: 0, Z: focal_length})

func rayColor(r *ray.Ray) *vec3.Vec3 {
	unit_direction := r.Direction.UnitVector()
	t := 0.5 * (unit_direction.Y + 1.0)
	v1 := vec3.Vec3{X: 1.0, Y: 1.0, Z: 1.0}
	v2 := vec3.Vec3{X: 0.5, Y: 0.7, Z: 1.0}
	return v1.Scale(1.0 - t).Plus(v2.Scale(t))
}

func main() {
	fmt.Printf("P3\n%d %d\n255\n", image_width, image_height)
	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < image_width; i++ {
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)
			// Direction is lower_left_corner + u*horizontal + v*vertical - origin
			r := ray.Ray{Origin: origin, Direction: *lower_left_corner.Plus(horizontal.Scale(u)).Plus(vertical.Scale(v)).Minus(&origin)}
			color := rayColor(&r)
			color.WriteColor(os.Stdout)
		}
	}
	fmt.Fprint(os.Stderr, "\nDone\n")

}
