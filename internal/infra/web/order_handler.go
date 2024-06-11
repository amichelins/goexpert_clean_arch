package web

import (
    "encoding/json"
    "net/http"

    "github.com/amichelins/goexpert_clean_arch/internal/dto"
    "github.com/amichelins/goexpert_clean_arch/internal/entity"
    "github.com/amichelins/goexpert_clean_arch/internal/usecase"
)

type WebOrderHandler struct {
    OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(
    OrderRepository entity.OrderRepositoryInterface,
) *WebOrderHandler {
    return &WebOrderHandler{
        OrderRepository: OrderRepository,
    }
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
    var OrderInput dto.OrderInputDTO

    err := json.NewDecoder(r.Body).Decode(&OrderInput)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository)

    OrderOuput, err := createOrder.Execute(OrderInput)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = json.NewEncoder(w).Encode(OrderOuput)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
    ListOrder := usecase.NewListOrderUseCase(h.OrderRepository)

    ListOrders, err := ListOrder.Execute()

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = json.NewEncoder(w).Encode(ListOrders)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
