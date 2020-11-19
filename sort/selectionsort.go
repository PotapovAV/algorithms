package sort

func SelectionSort(data []Data) []Data {
	for i := 0; i < len(data); i++ {
		minIdx := i
		for j := i + 1; j < len(data); j++ {
			if data[j].LessThan(data[minIdx]) {
				minIdx = j
			}
		}
		if minIdx != i {
			data[i], data[minIdx] = data[minIdx], data[i]
		}
	}
	return data
}
