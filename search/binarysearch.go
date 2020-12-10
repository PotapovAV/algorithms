package search

type BSData interface {
	LessThan(other interface{}) bool
	EqualsTo(other interface{}) bool
}

func BinarySearch(input []BSData, item BSData) (int, bool) {
	leftIdx := 0
	rightIdx := len(input) - 1
	return binarySearch(input, leftIdx, rightIdx, item)
}

func binarySearch(input []BSData, leftIdx, rightIdx int, item BSData) (int, bool) {
	if rightIdx-leftIdx == 0 {
		if input[leftIdx].EqualsTo(item) {
			return leftIdx, true
		}
		return -1, false
	}
	q := (rightIdx + leftIdx) / 2
	if input[q+1].EqualsTo(item) {
		return q + 1, true
	}
	if item.LessThan(input[q+1]) {
		return binarySearch(input, leftIdx, q, item)
	}
	return binarySearch(input, q+1, rightIdx, item)
}
