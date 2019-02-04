package main

import (
	"math"
)

func main() {
	ivory := Material{vec3{100, 100, 50}}
	//red := Material{vec3{200, 100, 100}}

	// Spheres in scene
	spheres := []Sphere{
		Sphere{vec3{0, 0, -16}, 5, ivory},
	}

	// Rest of the scene
	s := Scene{
		SphereToShape(spheres),
		[]Light{
			{vec3{15, 5, 2}, 1},
		},
		Camera{vec3{}, math.Pi / 2, 1024, 768},
	}

	s.render()
}
