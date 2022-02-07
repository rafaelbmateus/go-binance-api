package usecase

import "github.com/adshao/go-binance/v2"

func (c UsecaseService) ExchangeInfo() (*binance.ExchangeInfo, error) {
	exchangeInfo, err := c.Binance.NewExchangeInfoService().Do(c.Context)
	if err != nil {
		return nil, err
	}

	return exchangeInfo, nil
}
