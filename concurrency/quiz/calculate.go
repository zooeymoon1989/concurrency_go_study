package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println(calculatePi(50000000))
}

func calculatePi(n int) float64 {
	var squareTotal, circleTotal int
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if math.Pow(x, float64(2))+math.Pow(y, float64(2)) < 1 {
			circleTotal++
		}
		squareTotal++
	}
	value, _ := strconv.ParseFloat(fmt.Sprintf("%1.10f", float64(4 * circleTotal) / float64(squareTotal)), 64)
	return value
}
