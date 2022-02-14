package usecase

import (
	"github.com/adshao/go-binance/v2"
)

// CreateOrder
func (c UsecaseService) CreateOrder(symbol, quantity, price string) (*binance.CreateOrderResponse, error) {
	order, err := c.Binance.NewCreateOrderService().Symbol(symbol).
		Side(binance.SideTypeBuy).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantity).
		Price(price).Do(c.Context)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// GetOrder
func (c UsecaseService) GetOrder(orderID int64) (*binance.Order, error) {
	order, err := c.Binance.NewGetOrderService().OrderID(orderID).Do(c.Context)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (c UsecaseService) CancelOrder(orderID int64) error {
	_, err := c.Binance.NewCancelOrderService().OrderID(orderID).Do(c.Context)
	if err != nil {
		return err
	}

	return nil
}

func (c UsecaseService) ListOpenOrders() ([]*binance.Order, error) {
	openOrders, err := c.Binance.NewListOpenOrdersService().Do(c.Context)
	if err != nil {
		return nil, err
	}

	return openOrders, nil
}

func (c UsecaseService) ListOrders() ([]*binance.Order, error) {
	orders, err := c.Binance.NewListOrdersService().Do(c.Context)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
