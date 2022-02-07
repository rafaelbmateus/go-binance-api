package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func GetAccount(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler get account")

		acc, err := s.UsecaseService.GetAccount()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, acc)
	})
}
