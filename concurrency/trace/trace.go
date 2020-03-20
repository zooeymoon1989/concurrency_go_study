package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

const (
	output     = "out.png"
	width      = 2048
	height     = 2048
	numWorkers = 8
	cpuPprof   = "cpu.pprof"
)

func main() {

	trace.Start(os.Stdout)
	defer trace.Stop()

	f, err := os.Create(output)
	if err != nil {
		fmt.Printf("can't not open file err : %v", err.Error())
		return
	}

	img := createWorkers(width, height)

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}

}

func create(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}

	return m
}

func createPix(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	wg.Add(width * height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			go func(i, j int) {
				m.Set(i, j, pixel(i, j, width, height))
				wg.Done()
			}(i, j)

		}
	}

	wg.Wait()

	return m
}

func createRow(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	wg.Add(width * height)
	for i := 0; i < width; i++ {

		go func(i int) {
			for j := 0; j < height; j++ {
				m.Set(i, j, pixel(i, j, width, height))
				wg.Done()
			}
		}(i)

	}

	wg.Wait()

	return m
}

func createWorkers(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup
	c := make(chan int , width)
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func() {
			for p := range c {
				for j := 0 ; j < height ; j++ {
					m.Set(p, j, pixel(p, j, width, height))
				}

			}
			wg.Done()
		}()

	}
	for i := 0; i < width; i++ {
		c <- i
	}
	close(c)
	wg.Wait()

	return m
}

// createCol creates one goroutine per column.
func createCol(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			m.Set(i, j, pixel(i, j, width, height))
		}
	}
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
