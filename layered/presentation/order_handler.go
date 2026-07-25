package presentation

import (
	"fmt"

	"github.com/jordanmarta/software-architecture-labs/layered/business"
	"github.com/jordanmarta/software-architecture-labs/layered/model"
)

type OrderHandler struct {
	service business.OrderService
}

func NewOrderHandler(service business.OrderService) OrderHandler {
	return OrderHandler{
		service: service,
	}
}

func (h OrderHandler) CreateOrder(customer string, items []model.Item) {
	order := model.Order{
		Customer: customer,
		Items:    items,
	}

	createdOrder, err := h.service.CreateOrder(order)
	if err != nil {
		fmt.Println("error creating order:", err)
		return
	}

	fmt.Printf(
		"[Presentation] Order created for %s | total: %d cents\n",
		createdOrder.Customer,
		createdOrder.Total,
	)
}
