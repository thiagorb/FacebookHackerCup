package optimized

func CanReachRightmostLilypad(input string) bool {
	if len(input) < 3 {
		return false
	}

	betasCount := 0

	for i := 1; i < len(input); i++ {
		if input[i] == 'B' {
			betasCount++
		}
	}

	if len(input) == 3 {
		return betasCount == 1
	}

	return betasCount >= 2 && betasCount <= len(input)-2
}
