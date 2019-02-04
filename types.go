package main

// Scene ontaines all elements of a scene
type Scene struct {
	Objects []Shape
	Lights  []Light
	Cam     Camera
}

// Shape is an interface to graphics
type Shape interface {
	rayIntersect(orig, dir vec3) (hit bool, hitpos, normal vec3)
	Color() Material
}

// Sphere is a round object
type Sphere struct {
	Center vec3
	Radius float64
	Mat    Material
}

// Light is a source of light in the scene
type Light struct {
	Pos    vec3
	Intens float64
}

// Camera is a view point
type Camera struct {
	pos                vec3
	fov, width, height float64
}

type vec3 [3]float64

// Material is what an object is made of
type Material struct {
	color vec3
}
