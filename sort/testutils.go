package sort

type testType int64

func (s testType) LessThan(other interface{}) bool {
	switch other.(type) {
	case testType:
		if s < other.(testType) {
			return true
		}
	default:
		panic("Wrong data type provided")
	}
	return false
}

func (s testType) EqualsTo(other interface{}) bool {
	switch other.(type) {
	case testType:
		if s == other.(testType) {
			return true
		}
	default:
		panic("Wrong data type provided")
	}
	return false
}

func getDataSlice(inputData []testType) []Data {
	container := make([]Data, len(inputData))
	for idx, value := range inputData {
		container[idx] = testType(value)
	}
	return container
}
