package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // search for command line arguments
	match := false
	for _, arg := range args { // to search matching strings
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			match = true
			break
		}
	}
	if match {
		fmt.Println("Alert!!!")
	}
}
