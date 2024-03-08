package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Params := os.Args[1:]
	for i := 0; i < len(Params)-1; i++ {
		for j := i + 1; j < len(Params); j++ {
			if Params[j] < Params[i] {
				Params[i], Params[j] = Params[j], Params[i]
			}
		}
	}
	for k := 0; k < len(Params); k++ {
		for _, value := range Params[k] {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
