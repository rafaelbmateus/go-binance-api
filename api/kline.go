package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetKlines(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler get account")

		acc, err := s.UsecaseService.GetKlines(getSymbol(c), "1m")
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, acc)
	})
}
