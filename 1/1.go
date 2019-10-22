package main

import (
	"fmt"
	"math"
)

func update(x, z float64) float64 {
	return z - (z*z-x)/(2*z)
}

func isClose(x, z float64) bool {
	tolerance := 1.0E-7
	return math.Abs(z-x) < tolerance
}

func Sqrt(x float64) float64 {
	max_tries := 10
	z := x
	for i := 0; i < max_tries; i++ {
		z = update(x, z)
		if isClose(x, z*z) {
			return z
		}
	}
	return 0.0
}

func main() {
	fmt.Println(Sqrt(2))
}
