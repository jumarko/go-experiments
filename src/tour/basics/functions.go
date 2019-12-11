package main

import "fmt"

// could also be:
//   func add( x, y int) int {
func add( x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))
	// a function can return multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
