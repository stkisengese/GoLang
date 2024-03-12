package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point, newx int, newy int) {
	ptr.x = newx
	ptr.y = newy
}

func main() {
	points := point{x: 0, y: 1}

	setPoint(&points, 42, 21)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	printInt(points.x)

	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	printInt(points.y)

	z01.PrintRune('\n')
}

func printInt(n int) {
	if n < 0 { // handle the negative numbers
		z01.PrintRune('-')
		n = -n
	}
	// convert n to a string of digits
	var result []rune
	for n > 0 {
		digit := rune('0' + n%10)
		result = append(result, digit)
		n /= 10
	}
	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
		// print digits
		for _, digit := range result {
			z01.PrintRune(digit)
		}
	}
}
