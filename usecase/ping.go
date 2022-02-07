package usecase

func (c UsecaseService) Ping() error {
	err := c.Binance.NewPingService().Do(c.Context)
	if err != nil {
		return err
	}

	return nil
}
