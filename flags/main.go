package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		return
	}
	result := ""
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--insert", "-i":
			if i+1 < len(args) {
				result += args[i+1]
				i++
			}
		case "--order", "-o":
			if i+1 < len(args) {
				result = (result + args[i+1])
				i++
			}
		default:
			result += args[i]
		}
	}
	fmt.Println(result)
}
