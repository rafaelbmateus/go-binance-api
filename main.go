package main

import (
	"context"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/rafaelbmateus/go-binance-bot/api"
	"github.com/rs/zerolog"
)

var (
	Version = "no version provided"
	Commit  = "no commit hash provided"
)

func main() {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_API_SECRET")

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	log.Info().Msgf("starting with version %s and commit %s", Version, Commit)

	client := binance.NewClient(apiKey, secretKey)
	s := api.NewServer(context.Background(), &log, client)
	if err := s.Server().Run(os.Getenv("HOST")); err != nil {
		log.Fatal().Msg("error on run server")
	}
}
