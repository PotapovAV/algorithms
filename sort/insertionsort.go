package sort

func InsertionSort(data []Data) []Data {
	for i := 1; i < len(data); i++ {
		for j := i - 1; (j >= 0) && (data[j+1].LessThan(data[j])); j-- {
			data[j+1], data[j] = data[j], data[j+1]
		}
	}
	return data
}
