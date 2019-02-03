package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func castRay(orig, dir vec3, sps []shape) vec3 {
	for _, s := range sps {
		if s.rayIntersect(orig, dir) {
			return vec3{255, 0, 0}
		}
	}
	return vec3{0, 0, 0}
}

func render(sps []shape) {
	//var width, height float64 = 2880 / 4, 1800 / 4
	var width, height float64 = 1024, 768
	frame := make([]vec3, int(width*height))
	fov := 3.1415 / 2

	for i := 0.0; i < float64(height); i++ {
		for j := 0.0; j < float64(width); j++ {
			x := (2*(j+0.5)/float64(width) - 1) * math.Tan(fov/2) * float64(width/height)
			y := -(2*(i+0.5)/float64(height) - 1) * math.Tan(fov/2)
			frame[int(i*width+j)] = castRay(vec3{0, 0, 0}, vec3{x, y, -1}, sps)
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

	fmt.Fprintf(w, "P3\n%d %d\n255\n", int(width), int(height))
	for i := 0.0; i < height; i++ {
		for j := 0.0; j < width; j++ {
			for _, k := range frame[int(i*width+j)] {
				fmt.Fprintf(w, "%v ", int(k))
			}
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func main() {
	spheres := []Sphere{
		{vec3{-3, 0, -16}, 2},
		{vec3{12, 5, -25}, 4},
	}
	shapes := make([]shape, len(spheres))
	for i := range shapes {
		shapes[i] = spheres[i]
	}
	render(shapes)
}
