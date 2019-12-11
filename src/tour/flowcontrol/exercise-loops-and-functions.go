package main

// https://tour.golang.org/flowcontrol/8

import (
	"fmt"
)

func Abs(x float64) float64     {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Sqrt(x float64) float64 {
	guess := 1.0
	for {
		newGuess := guess - (guess * guess - x) / (2 * guess)
		if Abs(newGuess - guess) < 0.1 {
			return newGuess
		}
	}

}

func main() {
	fmt.Println(Sqrt(2))
}

