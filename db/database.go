package db

import (
	"github.com/rafaelbmateus/go-binance-bot/entity"
)

type Repository interface {
	GetStrategies() (*[]entity.Strategy, error)
	GetStrategy(id int) (*entity.Strategy, error)
	InsertStrategy(strategy *entity.Strategy) error
	UpdateStrategy(id int, strategy *entity.Strategy) (*entity.Strategy, error)
	DeleteStrategy(id int) error
}
