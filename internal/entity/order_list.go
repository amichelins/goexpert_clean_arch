package entity

import "errors"

type Order struct {
    ID         string
    Price      float64
    Tax        float64
    FinalPrice float64
}

// NewOrder Cria uma nova Order a partir dos parâmetros recebidos
//
// PARAMETERES
//
//     id string Id da order
//
//     price float64 Valor de preço
//
//     tax float64 Valor da Tax
//
// RETURN
//
//     *Order Order criada, só enviamos o endereço da memoria
//
//     error Erro inicializado ou nil em caso de sucesso
//
func NewOrder(id string, price float64, tax float64) (*Order, error) {
    // Criamos a order
    order := &Order{
        ID:    id,
        Price: price,
        Tax:   tax,
    }

    // Validamos que os dados sejam corretos
    err := order.IsValid()

    if err != nil {
        return nil, err
    }
    return order, nil
}

// IsValid Verifica se a Order tem dados validos segundo as regras especificadas, se não
//         gera um erro para cada situação encontrada
//
// PARAMETERES
//
// RETURN
//
//     error Erro inicializado ou nil em caso de sucesso
//
func (o *Order) IsValid() error {
    if o.ID == "" {
        return errors.New("ID invalido.")
    }

    if o.Price <= 0 {
        return errors.New("Price invalido. Price menor o igual a zero")
    }
    if o.Tax <= 0 {
        return errors.New("Tax invalido. Tax menor o igual a zero ")
    }

    return nil
}

// CalculateFinalPrice Calcula o preço final a partir dos dados da Order
//
// PARAMETERES
//
// RETURN
//
//     error Erro inicializado ou nil em caso de sucesso
//
func (o *Order) CalculateFinalPrice() error {
    // Calculamos o preço final
    o.FinalPrice = o.Price + o.Tax

    // Verificamos se a Order é valida
    err := o.IsValid()

    if err != nil {
        return err
    }
    return nil
}
