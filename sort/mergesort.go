package sort

import (
	"runtime"
	"sync"
)

func MergeSort(input, buffer []Data, p, r int) []Data {
	if p < r {
		q := (p + r) / 2
		MergeSort(input, buffer, p, q)
		MergeSort(input, buffer, q+1, r)
		merge(input, buffer, p, q, r)
	}
	return input
}

func merge(input, buffer []Data, p, q, r int) {
	leftIdx := p
	rightIdx := q + 1
	for bufferIdx := p; bufferIdx <= r; bufferIdx++ {
		if (leftIdx <= q) && (rightIdx <= r) {
			if input[leftIdx].LessThan(input[rightIdx]) {
				buffer[bufferIdx] = input[leftIdx]
				leftIdx++
			} else {
				buffer[bufferIdx] = input[rightIdx]
				rightIdx++
			}
		} else {
			if leftIdx > q {
				buffer[bufferIdx] = input[rightIdx]
				rightIdx++
			} else {
				buffer[bufferIdx] = input[leftIdx]
				leftIdx++
			}
		}
	}
	for bufferIdx := p; bufferIdx <= r; bufferIdx++ {
		input[bufferIdx] = buffer[bufferIdx]
	}
}

func ParallelMergeSort(input, buffer []Data, p, r int) []Data {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	if p < r {
		coroutinesAmount := numCPU
		if (r - p + 1) < numCPU {
			coroutinesAmount = r - p + 1
		}
		var wg sync.WaitGroup
		wg.Add(coroutinesAmount)

		q := (p + r) / coroutinesAmount
		pIdx := 0
		rIdx := q
		for coroutineNum := 1; coroutineNum <= coroutinesAmount; coroutineNum++ {
			go func(pIdx, rIdx int) {
				MergeSort(input, buffer, pIdx, rIdx)
				wg.Done()
			}(pIdx, rIdx)

			pIdx = min(pIdx+q+1, r)
			rIdx = min(rIdx+q+1, r)
		}
		wg.Wait()
		mergeParallelResult(input, buffer, p, r, coroutinesAmount)
	}
	return input
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mergeParallelResult(input, buffer []Data, p, r, coroutinesAmount int) {
	if coroutinesAmount > 1 {
		q := (p + r) / 2
		mergeParallelResult(input, buffer, p, q, coroutinesAmount/2)
		if coroutinesAmount%2 == 0 {
			mergeParallelResult(input, buffer, q+1, r, coroutinesAmount/2)
		} else {
			mergeParallelResult(input, buffer, q+1, r, coroutinesAmount/2+1)
		}
		merge(input, buffer, p, q, r)
	}
}
