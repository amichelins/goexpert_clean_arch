package service

import (
    "context"

    "github.com/amichelins/goexpert_clean_arch/internal/infra/grpc/pb"
    "github.com/amichelins/goexpert_clean_arch/internal/usecase"
)

type OrderService struct {
    pb.UnimplementedOrderServiceServer
    ListOrderUseCase usecase.ListOrderUseCase
}

func NewOrderService(listOrderUseCase usecase.ListOrderUseCase) *OrderService {
    return &OrderService{ListOrderUseCase: listOrderUseCase}
}

func (s *OrderService) ListOrders(context.Context, *pb.Blank) (*pb.OrderList, error) {
    LisOrder, err := s.ListOrderUseCase.Execute()

    if err != nil {
        return nil, err
    }

    PbOrderList := pb.OrderList{}

    for _, Order := range LisOrder {
        PbOrderList.Orders = append(PbOrderList.Orders, &pb.OrderResponse{Id: Order.ID, Price: float32(Order.Price), Tax: float32(Order.Tax), FinalPrice: float32(Order.FinalPrice)})
    }

    return &PbOrderList, nil
}
