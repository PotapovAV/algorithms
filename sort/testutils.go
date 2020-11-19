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

func getDataSlice(input []testType) []Data {
	container := make([]Data, len(input))
	for idx, value := range input {
		container[idx] = value
	}
	return container
}

func getQSDataSlice(input []testType) []QSData {
	container := make([]QSData, len(input))
	for idx, value := range input {
		container[idx] = value
	}
	return container
}
