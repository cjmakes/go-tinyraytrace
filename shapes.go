package main

type shape interface {
	rayIntersect(orig, dir vec3) bool
}

// Sphere is a round object
type Sphere struct {
	Center vec3
	Radius float64
}

// Orig is a point, dir is a vector
func (s Sphere) rayIntersect(orig, dir vec3) bool {
	oc := s.Center.Sub(orig)
	proj := dir.ScalarProduct(dir.DotProduct(oc) / dir.Magnitude())
	dist := oc.Sub(proj)

	// The case where the ray does not hit the sphere
	if dist.Magnitude() > s.Radius {
		return false
	}
	return true
}
