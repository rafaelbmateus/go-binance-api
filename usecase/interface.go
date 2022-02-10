package usecase

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/rafaelbmateus/go-binance-bot/db"
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
	Context    context.Context
	Logger     *zerolog.Logger
	Binance    *binance.Client
	Repository db.Repository
}

func NewUsecaseService(ctx context.Context, log *zerolog.Logger,
	binance *binance.Client, repository db.Repository) *UsecaseService {
	return &UsecaseService{
		Context:    ctx,
		Logger:     log,
		Binance:    binance,
		Repository: repository,
	}
}
