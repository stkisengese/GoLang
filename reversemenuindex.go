package piscine

func ReverseMenuIndex(menu []string) []string {
	reversedMenu := make([]string, len(menu))

	for i := len(menu) - 1; i >= 0; i-- {
		reversedMenu[len(menu)-1-i] = menu[i]
	}
	return reversedMenu
}
