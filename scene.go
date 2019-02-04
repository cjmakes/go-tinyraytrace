package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func sceneIntersect(orig, dir vec3, objects []Shape) (hit bool, hitpos, normal vec3, mat Material) {
	sphereDist := math.MaxFloat64
	for _, o := range objects {
		thit, thitpos, tnormal := o.rayIntersect(orig, dir)
		if thit && thitpos.Magnitude() < sphereDist {
			sphereDist = hitpos.Magnitude()
			hitpos = thitpos
			normal = tnormal
			mat = o.Color()
		}
	}
	hit = (sphereDist != math.MaxFloat64)
	if !hit {
		mat = Material{vec3{50, 75, 255}}
	}
	return
}

func castRay(orig, dir vec3, sps []Shape, lts []Light) Material {
	hit, hitpos, norm, mat := sceneIntersect(orig, dir, sps)
	if !hit {
		return mat
	}
	diffuseIntense := 0.0
	for _, l := range lts {
		lightDir := l.Pos.Sub(hitpos)
		lightDir = lightDir.Norm()
		if lightDir.DotProduct(norm) > 0 {
			diffuseIntense += l.Intens * lightDir.DotProduct(norm)
		}
	}
	m := Material{mat.color.ScalarProduct(diffuseIntense)}
	return m
}

func (scn *Scene) render() {
	frame := make([]vec3, int(scn.Cam.width*scn.Cam.height))
	// The ratio of half the BC / AB
	tan := math.Tan(scn.Cam.fov / 2)

	// Render Screen
	for i := 0.0; i < scn.Cam.height; i++ {
		for j := 0.0; j < scn.Cam.width; j++ {
			x := (2*(j+0.5)/scn.Cam.width - 1) * scn.Cam.width / scn.Cam.height
			y := -(2*(i+0.5)/scn.Cam.height - 1) * tan

			dir := vec3{x, y, -1}
			frame[int(i*scn.Cam.width+j)] =
				castRay(vec3{}, dir.Norm(), scn.Objects, scn.Lights).color
		}
	}

	// File operations
	outfile := os.Getenv("RAYOUT")
	_, err := os.Stat(outfile)
	if err == nil {
		os.Remove(outfile)
	}
	out, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Could not create file")
		return
	}
	defer out.Close()
	w := bufio.NewWriter(out)

	// Write to file
	fmt.Fprintf(w, "P3\n%d %d\n255\n", int(scn.Cam.width), int(scn.Cam.height))
	for i := 0.0; i < scn.Cam.height; i++ {
		for j := 0.0; j < scn.Cam.width; j++ {
			for _, k := range frame[int(i*scn.Cam.width+j)] {
				fmt.Fprintf(w, "%v ", int(k))
			}
		}
		w.WriteString("\n")
	}
	w.Flush()
}
