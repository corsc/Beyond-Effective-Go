package _2_not_leaking

type OrderManager struct {
	storage Storage
}

type Storage interface {
	LoadByID(int) ([]interface{}, error)
	Save(*Order) error
}

type Order struct {
	// implementation removed
}
