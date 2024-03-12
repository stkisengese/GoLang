package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("File name missing")
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		data, _ := os.ReadFile(args[1])
		fmt.Printf(string(data))
	}
}
