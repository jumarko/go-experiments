package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// or this way: https://tour.golang.org/flowcontrol/2
	// https://tour.golang.org/flowcontrol/3
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}


