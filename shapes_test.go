package main

import "testing"

func TestSphereRayIntersect(t *testing.T) {
	s := Sphere{vec3{0, 0, 5}, 2}

	tables := []struct {
		orig, dir vec3
		outcome   bool
	}{
		{vec3{0, 0, 0}, vec3{0, 0, 1}, true},
		{vec3{0, 0, 0}, vec3{10, 0, 1}, false},
		{vec3{0, 0, 0}, vec3{-10, 0, 1}, false},
		{vec3{0, 0, 0}, vec3{0, 10, 1}, false},
		{vec3{0, 0, 0}, vec3{0, -10, 1}, false},
	}

	for _, e := range tables {
		if s.rayIntersect(e.orig, e.dir) != e.outcome {
			t.Errorf("Orig %v Dir %v Expected %v", e.orig, e.dir, e.outcome)
		}
	}

}
