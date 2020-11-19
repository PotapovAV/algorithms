package sort

import "math"

func RadixSort(input []int64, digits int) []int64 {
	for i := 0; i < digits; i++ {
		input = radixSort(input, i)
	}
	return input
}

func radixSort(input []int64, digitPosition int) []int64 {
	output := make([]int64, len(input))
	buffer := make([]int64, 10)
	for _, item := range input {
		digit := item / int64(math.Pow10(digitPosition)) % int64(10)
		buffer[digit]++
	}
	for i := 1; i < len(buffer); i++ {
		buffer[i] += buffer[i-1]
	}
	for i := len(input) - 1; i >= 0; i-- {
		digit := input[i] / int64(math.Pow10(digitPosition)) % int64(10)
		output[buffer[digit]-1] = input[i]
		buffer[digit]--
	}
	return output
}
