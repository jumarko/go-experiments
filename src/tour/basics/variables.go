package main

import "fmt"

// https://tour.golang.org/basics/8
var c, python, java bool

// constants: https://tour.golang.org/basics/15
const MyPi = 3.14

func main() {
	// https://tour.golang.org/basics/9
	var i, j int = 1, 2
	// or you can use just := inside a function
	// https://tour.golang.org/basics/10
	k := 3
	fmt.Println(i, j, k, c, python, java)
	fmt.Println(MyPi)
}

