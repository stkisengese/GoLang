package piscine

func StrRev(s string) string {
	var result string /* empty string to store the reversed result */

	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i]) // append to result string
	}
	return result
}
