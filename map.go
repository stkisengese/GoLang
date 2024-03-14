package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a)) // slice of bool with same length as input slice
	for i, value := range a {
		result[i] = f(value) // store the result of f(value) in the result slice
	}
	return result // slice containing function results
}
