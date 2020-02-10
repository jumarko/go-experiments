package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	// try what happens when you use os.Args[0:]
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
