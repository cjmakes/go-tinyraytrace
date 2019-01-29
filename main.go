package main

import (
	"bufio"
	"fmt"
	"os"
)

func render() {
	width, height := 2880/4, 1800/4
	frame := make([][]float32, width*height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			frame[i*width+j] =
				[]float32{float32(i) / float32(height), float32(j) / float32(width), 0}
		}
	}

	out, _ := os.Create("out.ppm")
	defer out.Close()
	w := bufio.NewWriter(out)

	fmt.Fprintf(w, "P3\n%d %d\n255\n", width, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			for _, k := range frame[(i*width)+j] {
				fmt.Fprintf(w, "%v ", int(255*k))
			}
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func main() {
	render()
}
