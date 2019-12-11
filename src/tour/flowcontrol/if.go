package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if can also have a short "pre-condition" statement: https://tour.golang.org/flowcontrol/6
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	// https://tour.golang.org/flowcontrol/7
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

}

