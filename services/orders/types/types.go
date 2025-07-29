package types

import (
	"context"

	orders "github.com/empaid/ecom/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}
