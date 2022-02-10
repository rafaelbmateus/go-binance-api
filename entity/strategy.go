package entity

type Strategy struct {
	Symbol   string `json:"symbol"`
	Amount   string `json:"amount"`
	PriceLow string `json:"priceLow"`
	PriceHi  string `json:"priceHi"`
}
