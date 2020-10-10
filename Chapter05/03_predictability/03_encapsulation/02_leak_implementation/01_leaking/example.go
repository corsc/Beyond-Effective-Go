package _1_leaking

type OrderManager struct {
	storage Storage
}

type Storage interface {
	DoQuery(q string, args ...interface{}) ([]interface{}, error)
}
