package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testType struct {
	key int
}

func (t testType) LessThan(other interface{}) bool {
	switch other.(type) {
	case testType:
		if t.key < other.(testType).key {
			return true
		}
	}
	return false
}

func (t testType) EqualsTo(other interface{}) bool {
	switch other.(type) {
	case testType:
		if t.key == other.(testType).key {
			return true
		}
	}
	return false
}

func getBSData(input []int) []BSData {
	data := make([]BSData, len(input))
	for i := 0; i < len(input); i++ {
		data[i] = testType{input[i]}
	}
	return data
}

func TestBinarySearch(t *testing.T) {
	input := []int{1, 3, 5, 7, 8}
	data := getBSData(input)

	idx, isFound := BinarySearch(data, testType{key: 1})
	assert.Equal(t, 0, idx, "invalid index of item returned")
	assert.Equal(t, true, isFound, "invalid status of searching returned")

	idx, isFound = BinarySearch(data, testType{key: 8})
	assert.Equal(t, 4, idx, "invalid index of item returned")
	assert.Equal(t, true, isFound, "invalid status of searching returned")

	idx, isFound = BinarySearch(data, testType{key: 6})
	assert.Equal(t, -1, idx, "invalid index of item returned")
	assert.Equal(t, false, isFound, "invalid status of searching returned")
}
