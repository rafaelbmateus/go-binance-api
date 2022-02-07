package usecase

import (
	"context"

	"github.com/adshao/go-binance/v2"
)

func (c UsecaseService) GetAccount() (*binance.Account, error) {
	acc, err := c.Binance.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	return acc, nil
}
