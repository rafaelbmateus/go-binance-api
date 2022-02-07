package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CreateOrderReq struct {
	quantity string
	price    string
}

func CreateOrder(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler create order")

		var req CreateOrderReq
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		order, err := s.UsecaseService.CreateOrder(getSymbol(c), req.quantity, req.price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, order)
	})
}

func GetOrder(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler get order")

		orderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		acc, err := s.UsecaseService.GetOrder(getSymbol(c), orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, acc)
	})
}

func ListOrders(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler list of orders")

		orders, err := s.UsecaseService.ListOpenOrders(getSymbol(c))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, orders)
	})
}

func CancelOrder(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler cancel order")

		orderID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		}

		err = s.UsecaseService.CancelOrder(getSymbol(c), orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, "order-canceled")
	})
}

func getSymbol(c *gin.Context) string {
	return c.Param("symbol")
}
