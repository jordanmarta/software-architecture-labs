package business

import (
	"errors"

	"github.com/jordanmarta/software-architecture-labs/layered/model"
	"github.com/jordanmarta/software-architecture-labs/layered/persistence"
)

type OrderService struct {
	repository persistence.OrderRepository
}

func NewOrderService(repository persistence.OrderRepository) OrderService {
	return OrderService{
		repository: repository,
	}
}

func (s OrderService) CreateOrder(order model.Order) (model.Order, error) {
	if len(order.Items) == 0 {
		return model.Order{}, errors.New("order must have at least one item")
	}
	total := 0

	for _, item := range order.Items {
		total += item.Price * item.Quantity
	}

	order.Total = total

	err := s.repository.Save(order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}
