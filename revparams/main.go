package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Params := os.Args[1:]
	for i := len(Params) - 1; i >= 0; i-- {
		arguments := Params[i]
		for _, value := range arguments {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
