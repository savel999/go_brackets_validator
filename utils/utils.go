package utils

func InArray(str string, slice []string) int {
	for i, item := range slice {
		if item == str {
			return i
		}
	}
	return -1
}
