package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	initial := getDataSlice([]testType{66, 8, 7, 44})
	expected := getDataSlice([]testType{7, 8, 44, 66})
	sorted := InsertionSort(initial)
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}
