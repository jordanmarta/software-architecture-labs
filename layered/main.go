package main

import (
	"github.com/jordanmarta/software-architecture-labs/layered/business"
	"github.com/jordanmarta/software-architecture-labs/layered/model"
	"github.com/jordanmarta/software-architecture-labs/layered/persistence"
	"github.com/jordanmarta/software-architecture-labs/layered/presentation"
)

func main() {
	repository := persistence.OrderRepository{}
	service := business.NewOrderService(repository)
	handler := presentation.NewOrderHandler(service)

	items := []model.Item{
		{
			Name:     "Keyboard",
			Price:    15000,
			Quantity: 1,
		},
		{
			Name:     "Mouse",
			Price:    8000,
			Quantity: 2,
		},
	}

	handler.CreateOrder("Jordan", items)
}
