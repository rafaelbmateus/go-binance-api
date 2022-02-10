package api

import (
	"context"

	"github.com/adshao/go-binance/v2"
	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-binance-bot/db"
	"github.com/rafaelbmateus/go-binance-bot/usecase"
	"github.com/rs/zerolog"
)

type API struct {
	Router *gin.Engine
}

type Server struct {
	UsecaseService *usecase.UsecaseService
}

func NewServer(ctx context.Context, log *zerolog.Logger, binance *binance.Client, repository db.Repository) *Server {
	return &Server{
		UsecaseService: usecase.NewUsecaseService(context.Background(), log, binance, repository),
	}
}

func (s *Server) Server() *gin.Engine {
	r := gin.Default()
	r.GET("/", Home)
	r.GET("/ping", Ping(s))
	r.GET("/accounts/me", GetAccount(s))
	r.GET("/exchanges", ExchangeInfo(s))
	r.GET("/symbols/:symbol/orders", ListOrders(s))
	r.POST("/symbols/:symbol/orders", CreateOrder(s))
	r.GET("/symbols/:symbol/orders/:id", GetOrder(s))
	r.DELETE("/symbols/:symbol/orders/:id", CancelOrder(s))
	r.GET("/symbols/:symbol/klines", GetKlines(s))
	r.GET("/strategies", ListStrategies(s))
	r.POST("/strategies", CreateStrategy(s))
	return r
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{"content": "Welcome to go-binance-api"})
}
