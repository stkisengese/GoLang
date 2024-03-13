package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 { // check if arguments are provided
		data, err := ioutil.ReadAll(os.Stdin) // read from standard input
		if err != nil {
			z01.PrintRune('E') // print error rune
			z01.PrintRune('\n')
			os.Exit(1)
		}
		for _, b := range data { // loop each byte in data (assuming ASCII characters)
			z01.PrintRune(rune(b))
		}
	} else {
		for _, filename := range os.Args[1:] { // loop thro each argument(file)
			data, err := ioutil.ReadFile(filename) // open the file
			if err != nil {                        // print error
				for _, r := range []rune("Error opening file " + filename + ": ") {
					z01.PrintRune(r)
				}
				z01.PrintRune('\n')
				continue
			}
			for _, b := range data { // loop each byte in data (assuming ASCII characters)
				z01.PrintRune(rune(b))
			}
		}
	}
}
