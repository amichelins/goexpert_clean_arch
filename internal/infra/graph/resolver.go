package graph

import (
    "github.com/amichelins/goexpert_clean_arch/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
    ListOrderUseCase usecase.ListOrderUseCase
}
