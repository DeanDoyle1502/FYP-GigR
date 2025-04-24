package utils

func NormalizeUserIDs(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
