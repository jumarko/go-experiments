package main

import (
	"bufio" // to read the text
	"fmt"   // to print formatted output
	"io"    // io.Reader interface
	"os"    // to use OS resources
)

func main() {
	fmt.Println(count(os.Stdin))
}

// Counts the number of words/lines in the input
func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	wc := 0
	// ignoring errors here, but in your code you should always check for errors!
	for scanner.Scan() {
		wc++
	}
	return wc
}
