package main

import "math"

// Color returns the material of the sphere
func (s Sphere) Color() Material {
	return s.Mat
}

// orig is the origin of the ray
// dir is the unit vector
// Returns if the ray hit, and if so, what the point of intersection is
func (s Sphere) rayIntersect(orig, dir vec3) (hit bool, hitpos, normal vec3) {
	dir = dir.Norm()
	oc := s.Center.Sub(orig)
	OQ := orig.Add(dir.ScalarProduct(dir.DotProduct(oc)))
	CQ := oc.Sub(OQ)
	if CQ.Magnitude() > s.Radius {
		return false, vec3{}, vec3{}
	}
	lenPQ := math.Sqrt(s.Radius*s.Radius - CQ.Magnitude()*CQ.Magnitude())
	lenOP := OQ.Magnitude() - lenPQ
	hitpos = orig.Add(dir.ScalarProduct(lenOP))
	norm := hitpos.Sub(oc)

	return true, hitpos, norm.Norm()
}

// SphereToShape converts a slice of spheres to a slice of shapes
func SphereToShape(sps []Sphere) []Shape {
	shs := make([]Shape, len(sps))
	for i, s := range sps {
		shs[i] = s
	}
	return shs
}
