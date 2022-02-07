package usecase

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/rs/zerolog"
)

type Binance interface {
	Setup()
	CreateOrder()
	GetOrder()
	CancelOrder()
	ListOpenOrders()
	ListOrders()
	ListTickerPrices()
	ShowDepth()
	ListKlines()
	ListAggregateTrades()
	GetAccount() *binance.Account
	StartUserStream()
}

type UsecaseService struct {
	Context context.Context
	Logger  *zerolog.Logger
	Binance *binance.Client
}

func NewUsecaseService(ctx context.Context, log *zerolog.Logger, binance *binance.Client) *UsecaseService {
	return &UsecaseService{
		Context: ctx,
		Logger:  log,
		Binance: binance,
	}
}
