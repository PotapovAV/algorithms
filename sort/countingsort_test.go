package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingSort(t *testing.T) {
	initial := []int64{66, 7, 8, 44, 5, 3, 90, 40}
	expected := []int64{3, 5, 7, 8, 40, 44, 66, 90}
	sorted := CountingSort(initial, max(initial))
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}

func TestCountingSortInPlace(t *testing.T) {
	initial := []int64{66, 7, 8, 44, 5, 3, 90, 40}
	expected := []int64{3, 5, 7, 8, 40, 44, 66, 90}
	sorted := CountingSort(initial, max(initial))
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}

func max(input []int64) int64 {
	max := input[0]
	for _, item := range input {
		if item > max {
			max = item
		}
	}
	return max
}
