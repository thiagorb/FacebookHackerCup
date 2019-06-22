package brute

func isGoal(state []byte) bool {
	return state[len(state) - 1] == 'A'
}

func leap(state []byte, fromIndex int, toIndex int) []byte {
	newState := make([]byte, len(state))
	copy(newState, state)
	frog := state[fromIndex]
	newState[fromIndex] = '.'
	newState[toIndex] = frog
	return newState
}

func alphaCanLeapLeft(state []byte, alphaIndex int) (bool, int) {
	if alphaIndex <= 1 || state[alphaIndex - 1] == '.' {
		return false, 0
	}

	for i := alphaIndex - 2; i >= 0; i-- {
		if state[i] == '.' {
			return true, i
		}
	}

	return false, 0
}

func alphaCanLeapRight(state []byte, alphaIndex int) (bool, int) {
	if alphaIndex >= len(state) - 2 || state[alphaIndex + 1] == '.' {
		return false, 0
	}

	for i := alphaIndex + 2; i < len(state); i++ {
		if state[i] == '.' {
			return true, i
		}
	}

	return false, 0
}

func CanReachRightmostLilypad(input string) bool {
	state := []byte(input)
	visitedStates := map[string]bool{}

	var _canReachRightmostLilypad func(state []byte) bool
	_canReachRightmostLilypad = func(state []byte) bool {
		if visitedStates[string(state)] {
			return false
		}

		visitedStates[string(state)] = true

		if isGoal(state) {
			return true
		}

		for i := 0; i < len(state); i++ {
			switch state[i] {
			case 'A':
				canLeapLeft, toIndex := alphaCanLeapLeft(state, i)
				if canLeapLeft {
					if _canReachRightmostLilypad(leap(state, i, toIndex)) {
						return true
					}
				}

				canLeapRight, toIndex := alphaCanLeapRight(state, i)
				if canLeapRight {
					if _canReachRightmostLilypad(leap(state, i, toIndex)) {
						return true
					}
				}
			case 'B':
				if i > 0 && state[i-1] == '.' {
					if _canReachRightmostLilypad(leap(state, i, i-1)) {
						return true
					}
				}

				if i < len(state)-1 && state[i+1] == '.' {
					if _canReachRightmostLilypad(leap(state, i, i+1)) {
						return true
					}
				}
			}
		}

		return false
	}

	return _canReachRightmostLilypad(state)
}
