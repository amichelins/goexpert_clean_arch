package entity

type OrderRepositoryInterface interface {
    List() (*[]Order, error)
    Save(order *Order) error
}
