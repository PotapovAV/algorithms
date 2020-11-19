package sort

func SelectionSort(input []Data) []Data {
	for i := 0; i < len(input); i++ {
		minIdx := i
		for j := i + 1; j < len(input); j++ {
			if input[j].LessThan(input[minIdx]) {
				minIdx = j
			}
		}
		if minIdx != i {
			input[i], input[minIdx] = input[minIdx], input[i]
		}
	}
	return input
}
