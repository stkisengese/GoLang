package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var preptime int
	switch order {
	case "burger":
		preptime = 15
	case "chips":
		preptime = 10
	case "nuggets":
		preptime = 12
	default:
		preptime = 404
	}
	return preptime
}
