package validate

func OutsideSelectionRange(input int, maxRange int) bool {
	if input <= maxRange && input > 0 {
		return false
	} else {
		return true
	}
}
