package postgres

import (
	"fmt"

	"github.com/rafaelbmateus/go-binance-bot/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(host, user, password, name, port string) (*Postgres, error) {
	db, err := connect(host, user, password, name, port)
	if err != nil {
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		DB: db,
	}, nil
}

func (db *Postgres) GetStrategies() (*[]entity.Strategy, error) {
	var strategies *[]entity.Strategy
	if result := db.DB.Find(&strategies); result.Error != nil {
		return nil, result.Error
	}

	return strategies, nil
}

func (db *Postgres) GetStrategy(id int) (*entity.Strategy, error) {
	var strategy *entity.Strategy
	if result := db.DB.First(&strategy, id); result.Error != nil {
		return nil, result.Error
	}

	return strategy, nil
}

func (db *Postgres) InsertStrategy(strategy *entity.Strategy) error {
	if result := db.DB.Create(&strategy); result.Error != nil {
		return result.Error
	}

	return nil
}

func (db *Postgres) UpdateStrategy(id int, strategy *entity.Strategy) (*entity.Strategy, error) {
	if result := db.DB.First(strategy, id); result.Error != nil {
		return nil, result.Error
	}

	return strategy, nil
}

func (db *Postgres) DeleteStrategy(id int) error {
	var strategy entity.Strategy
	if result := db.DB.First(&strategy, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func connect(host, user, password, name, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Strategy{})
}
