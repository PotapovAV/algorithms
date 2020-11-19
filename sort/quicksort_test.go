package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	initial := getQSDataSlice([]testType{66, 8, 7, 44})
	expected := getQSDataSlice([]testType{7, 8, 44, 66})
	sorted := QuickSort(initial, 0, len(initial)-1)
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}
