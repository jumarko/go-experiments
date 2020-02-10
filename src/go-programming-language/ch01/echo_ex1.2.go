// Ex. 1.2 ,p. 8:
// Modify the echo program to print the index and value of each of its arguments, one per line
package main

import (
	"fmt"
	"os"
	"strconv"
)

func printIndex1() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(i)
		fmt.Println(". " + os.Args[i])
	}
}

func printIndex2() {
	for i, arg := range os.Args[1:] {
		fmt.Print(i + 1)
		fmt.Println(". " + arg)
	}
}

func printIndex3() {
	for i, arg := range os.Args[1:] {
		// need to convert the int to string otherwise you get a type error
		fmt.Println(strconv.Itoa(i + 1) + ". " + arg)
	}
}


func main() {
	fmt.Println("-------------")
	printIndex1()
	fmt.Println("-------------")
	printIndex2()
	fmt.Println("-------------")
	printIndex3()
	fmt.Println("-------------")
}
