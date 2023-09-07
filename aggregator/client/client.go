package client

import (
	"context"

	"github.com/Jimbo8702/toll-calc/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
	GetInvoice(context.Context, int) (*types.Invoice, error)
}