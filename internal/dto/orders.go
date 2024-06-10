package dto

// O PKG dto, contem todos os dtos
// Como os uses cases utilizam exatamente os mesmo, eles são criados uma vez
type OrderInputDTO struct {
    ID    string  `json:"id"`
    Price float64 `json:"price"`
    Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
    ID         string  `json:"id"`
    Price      float64 `json:"price"`
    Tax        float64 `json:"tax"`
    FinalPrice float64 `json:"final_price"`
}
