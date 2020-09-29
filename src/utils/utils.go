package utils

// Find takes a slice of integers and looks for an element in it.
func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
