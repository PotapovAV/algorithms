package sort

type Data interface {
	LessThan(other interface{}) bool
}
