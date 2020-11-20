package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	initial := getHeapDataSlice([]testType{66, 8, 7, 44})
	expected := getHeapDataSlice([]testType{7, 8, 44, 66})
	sorted := HeapSort(initial)
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}
