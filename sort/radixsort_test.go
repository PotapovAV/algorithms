package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadixSort(t *testing.T) {
	initial := []int64{66, 7, 8, 44, 5, 3, 90, 40}
	expected := []int64{3, 5, 7, 8, 40, 44, 66, 90}
	sorted := RadixSort(initial, 2)
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}
