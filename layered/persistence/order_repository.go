package persistence

import (
	"fmt"

	"github.com/jordanmarta/software-architecture-labs/layered/model"
)

type OrderRepository struct{}

func (r OrderRepository) Save(order model.Order) error {
	fmt.Printf(
		"[Persistence] Saving order for %s | total: %d cents\n",
		order.Customer,
		order.Total,
	)

	return nil
}
