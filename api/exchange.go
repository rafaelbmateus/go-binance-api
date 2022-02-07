package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func ExchangeInfo(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler exchange info")

		symbols, err := s.UsecaseService.ExchangeInfo()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, symbols)
	})
}
