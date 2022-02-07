package usecase

import "github.com/adshao/go-binance/v2"

func (c UsecaseService) GetKlines(symbol, interval string) ([]*binance.Kline, error) {
	klines, err := c.Binance.NewKlinesService().Symbol(symbol).Interval(interval).Do(c.Context)
	if err != nil {
		return nil, err
	}

	return klines, nil
}
