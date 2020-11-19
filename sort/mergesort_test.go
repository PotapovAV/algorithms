package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	initial := getDataSlice([]testType{66, 8, 7, 44})
	expected := getDataSlice([]testType{7, 8, 44, 66})
	buffer := make([]Data, len(initial))
	sorted := MergeSort(initial, buffer, 0, len(initial)-1)
	assert.Equal(t, expected, sorted, "Sorted result do not match an expected one")
}

func TestParallelMergeSort(t *testing.T) {
	initial := getDataSlice([]testType{66, 8, 7, 44})
	expected := getDataSlice([]testType{7, 8, 44, 66})
	buffer := make([]Data, len(initial))
	sorted := ParallelMergeSort(initial, buffer, 0, len(initial)-1)
	assert.Equal(t, expected, sorted, "Sorted result do not match an expected one")
}
