package usecase

import "github.com/rafaelbmateus/go-binance-bot/entity"

func (c UsecaseService) ListStrategies() (*[]entity.Strategy, error) {
	return c.Repository.GetStrategies()
}

func (c UsecaseService) CreateStrategy(strategy *entity.Strategy) error {
	return c.Repository.InsertStrategy(strategy)
}

func (c UsecaseService) GetStrategy(id int) (*entity.Strategy, error) {
	return c.Repository.GetStrategy(id)
}

func (c UsecaseService) UpdateStrategy(id int, strategy *entity.Strategy) (*entity.Strategy, error) {
	return c.Repository.UpdateStrategy(id, strategy)
}

func (c UsecaseService) DeleteStrategy(id int) error {
	return c.Repository.DeleteStrategy(id)
}
