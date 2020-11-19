package sort

func CountingSort(input []int64, maxValue int64) []int64 {
	output := make([]int64, len(input))
	buffer := make([]int64, maxValue+1)
	for _, item := range input {
		buffer[item]++
	}
	for i := 1; i < len(buffer); i++ {
		buffer[i] += buffer[i-1]
	}
	for i := len(input) - 1; i >= 0; i-- {
		output[buffer[input[i]]-1] = input[i]
		buffer[input[i]]--
	}
	return output
}

func CountingSortInPlace(input []int64, maxValue int64) []int64 {
	buffer := make([]int64, maxValue+1)
	for _, item := range input {
		buffer[item]++
	}
	for j, i := int64(0), int64(0); j < int64(len(buffer)); j++ {
		if buffer[j] != 0 {
			amount := buffer[j]
			for amount > 0 {
				input[i] = j
				amount--
				i++
			}
		}
	}
	return input
}
