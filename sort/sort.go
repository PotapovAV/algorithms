package sort

type Data interface {
	LessThan(other interface{}) bool
}

type QSData interface {
	Data
	EqualsTo(other interface{}) bool
}
