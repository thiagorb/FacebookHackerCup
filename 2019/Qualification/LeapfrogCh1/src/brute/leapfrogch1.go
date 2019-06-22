package brute

func isGoal(state []byte) bool {
	return len(state) == 0
}

func alphaPositionAfterLeap(state []byte) (bool, int) {
	if state[0] == '.' {
		return false, 0
	}

	for i := 0; i < len(state); i++ {
		if state[i] == '.' {
			return true, i
		}
	}

	return false, 0
}

func CanReachRightmostLilypad(input string) bool {
	state := []byte(input)[1:]
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

		alphaCanLeap, alphaPositionAfterLeap := alphaPositionAfterLeap(state)
		if alphaCanLeap && _canReachRightmostLilypad(state[alphaPositionAfterLeap+1:]) {
			return true
		}

		for i := 1; i < len(state); i++ {
			if state[i] != 'B' {
				continue
			}

			if state[i-1] == '.' {
				newState := make([]byte, len(state))
				copy(newState, state)
				newState[i-1] = 'B'
				newState[i] = '.'
				if _canReachRightmostLilypad(newState) {
					return true
				}
			}

			if i < len(state)-1 && state[i+1] == '.' {
				newState := make([]byte, len(state))
				copy(newState, state)
				newState[i+1] = 'B'
				newState[i] = '.'
				if _canReachRightmostLilypad(newState) {
					return true
				}
			}
		}

		return false
	}

	return _canReachRightmostLilypad(state)
}
