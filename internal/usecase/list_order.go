package usecase

import (
    "github.com/amichelins/goexpert_clean_arch/internal/dto"
    "github.com/amichelins/goexpert_clean_arch/internal/entity"
)

type ListOrderUseCase struct {
    OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
    return &ListOrderUseCase{OrderRepository: OrderRepository}
}

func (c *ListOrderUseCase) Execute() ([]dto.OrderOutputDTO, error) {
    var OrdersOutput []dto.OrderOutputDTO

    Orders, err := c.OrderRepository.List()

    if err != nil {
        return OrdersOutput, err
    }

    for _, Order := range *Orders {
        OrdersOutput = append(OrdersOutput, dto.OrderOutputDTO{ID: Order.ID, Price: Order.Price, Tax: Order.Tax, FinalPrice: Order.FinalPrice})
    }

    return OrdersOutput, nil
}
