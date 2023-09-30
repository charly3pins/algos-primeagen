package search

func BinarySearch(hayStack []int, needle int) bool {
	low := 0
	high := len(hayStack)

	for low < high {
		mid := (low + high) / 2
		val := hayStack[mid]
		if val == needle {
			return true
		} else if val > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}
