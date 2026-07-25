package model

type Item struct {
	Name     string
	Price    int // valor em centavos
	Quantity int
}

type Order struct {
	Customer string
	Items    []Item
	Total    int
}
