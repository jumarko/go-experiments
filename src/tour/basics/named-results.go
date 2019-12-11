package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// this returns "named" return variables x, y - aka "naked return"
	return
}

func main() {
	fmt.Println(split(17))
}

