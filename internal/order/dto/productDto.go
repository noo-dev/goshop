package dto

type ProductDto struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
