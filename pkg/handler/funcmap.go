package handler

func Whole(a int, b int) bool {
	if a%b == 0 {
		return true
	} else if a%b == 1 {
		return false
	}
	return true
}
func Decimal(a float64) float64 {
	return a / 100
}
