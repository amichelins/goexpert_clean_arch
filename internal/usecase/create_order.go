package usecase

import (
    "github.com/amichelins/goexpert_clean_arch/internal/dto"
    "github.com/amichelins/goexpert_clean_arch/internal/entity"
)

type CreateOrderUseCase struct {
    OrderRepository entity.OrderRepositoryInterface
}

func NewCreateOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *CreateOrderUseCase {
    return &CreateOrderUseCase{OrderRepository: OrderRepository}
}

func (c *CreateOrderUseCase) Execute(input dto.OrderInputDTO) (dto.OrderOutputDTO, error) {
    order := entity.Order{
        ID:    input.ID,
        Price: input.Price,
        Tax:   input.Tax,
    }
    order.CalculateFinalPrice()

    if err := c.OrderRepository.Save(&order); err != nil {
        return dto.OrderOutputDTO{}, err
    }

    dto := dto.OrderOutputDTO{
        ID:         order.ID,
        Price:      order.Price,
        Tax:        order.Tax,
        FinalPrice: order.Price + order.Tax,
    }

    return dto, nil
}
