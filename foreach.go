package piscine

import "fmt"

func ForEach(f func(int), a []int) {
	for _, element := range a {
		f(element)
	}
}

func PrintNbr(n int) {
	fmt.Print(n)
}
