package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	uppercase := false
	if len(args) > 0 && args[0] == "--upper" {
		uppercase = true
		args = args[1:]
	}
	if len(args) == 0 {
		return
	}
	for _, arg := range args {
		num := 0
		for _, ch := range arg {
			if ch < '0' {
				z01.PrintRune(' ')
				continue
			}
			num = num*10 + int(ch-'0')
		}
		if num >= 1 && num <= 26 {
			if uppercase {
				z01.PrintRune('A' + rune(num) - 1)
			} else {
				z01.PrintRune('a' + rune(num) - 1)
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
