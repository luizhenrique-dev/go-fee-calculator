package entity

// Qualquer struct que existir que implemente os métodos da interface OrderRepositoryInterface, será considerado um OrderRepository
type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}
