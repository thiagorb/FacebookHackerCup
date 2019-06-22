package mrx

func CalculateMinimumModifications(e Expression) int {
	s := GetReduced(&e)
	if s == Symbol0 || s == Symbol1 {
		return 0
	} else {
		return 1
	}
}
