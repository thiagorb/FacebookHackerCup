package optimized

func CanReachRightmostLilypad(input string) bool {
	betasCount := 0

	for i := 1; i < len(input); i++ {
		if input[i] == 'B' {
			betasCount++
		}
	}

	return betasCount > (len(input) - 2) / 2 && betasCount <= len(input) - 2
}
