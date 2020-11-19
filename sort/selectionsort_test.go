package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	initial := getDataSlice([]testType{66, 8, 7, 44})
	expected := getDataSlice([]testType{7, 8, 44, 66})
	sorted := SelectionSort(initial)
	assert.Equal(t, expected, sorted, "Sorted result do not match an expected one")
}
