package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketSort(t *testing.T) {
	initial := []float64{0.5, 0.1, 0.4, 0.7}
	expected := []float64{0.1, 0.4, 0.5, 0.7}
	sorted := BucketSort(initial)
	assert.Equal(t, expected, sorted, "Sorted result does not match an expected one")
}
