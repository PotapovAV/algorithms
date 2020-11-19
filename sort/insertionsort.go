package sort

func InsertionSort(input []Data) []Data {
	for i := 1; i < len(input); i++ {
		for j := i - 1; (j >= 0) && (input[j+1].LessThan(input[j])); j-- {
			input[j+1], input[j] = input[j], input[j+1]
		}
	}
	return input
}
