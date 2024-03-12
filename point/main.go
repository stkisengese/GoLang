package main

import "github.com/01-edu/z01"

var point string = "x = 42, y = 21"

func main() {
	for _, xy := range point {
		z01.PrintRune(xy)
	}
	z01.PrintRune('\n')
}
