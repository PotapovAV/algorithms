package sort

type BSData float64

func (d BSData) LessThan(other interface{}) bool {
	switch other.(type) {
	case BSData:
		if d < other.(BSData) {
			return true
		}
	default:
		panic("Wrong data type provided")
	}
	return false
}

func BucketSort(input []float64) []float64 {
	devider := 4
	buckets := make([][]Data, devider)
	for i := 0; i < len(input); i++ {
		bucketIdx := int(float64(devider) * input[i])
		buckets[bucketIdx] = append(buckets[bucketIdx], Data(BSData(input[i])))
	}
	for i, j := 0, 0; i < devider; i++ {
		bucket := buckets[i]
		InsertionSort(bucket)
		for _, item := range bucket {
			input[j] = float64(item.(BSData))
			j++
		}
	}
	return input
}
