package sort

import "github.com/PotapovAV/data-structures/heap"

func HeapSort(input []heap.Data) []heap.Data {
	if len(input) > 0 {
		heap := heap.BuildHeap(input, 2)
		for i := len(input) - 1; i >= 0; i-- {
			input[i] = heap.ExtractMax()
		}
	}
	return input
}
