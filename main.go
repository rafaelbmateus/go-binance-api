package main

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/rafaelbmateus/go-binance-bot/api"
	"github.com/rafaelbmateus/go-binance-bot/db/postgres"
	"github.com/rs/zerolog"
)

var (
	Version = "no version provided"
	Commit  = "no commit hash provided"
)

func main() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_API_SECRET")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Info().Msgf("starting with version %s and commit %s", Version, Commit)

	binance := binance.NewClient(apiKey, secretKey)
	database, err := postgres.NewPostgres(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatal().Msgf("error on run server %v", err)
		return
	}

	s := api.NewServer(context.Background(), &log, binance, database)
	if err := s.Server().Run(os.Getenv("HOST")); err != nil {
		log.Fatal().Msgf("error on run server %v", err)
	}
}
