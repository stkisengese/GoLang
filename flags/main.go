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
				result = orderString(result + args[i+1])
				i++
			}
		default:
			result += args[i]
		}
	}
	fmt.Println(result)
}
func orderString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}
