package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafaelbmateus/go-binance-bot/entity"
	"github.com/rs/zerolog/log"
)

func ListStrategies(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler list strategies")

		strategies, err := s.UsecaseService.ListStrategies()
		if err != nil {
			log.Info().Msgf("error on list strategies %v", err.Error())
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, strategies)
	})
}

func GetStrategy(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler get strategy")

		strategy, err := s.UsecaseService.ListStrategies()
		if err != nil {
			log.Info().Msgf("error on list strategies %v", err.Error())
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, strategy)
	})
}

func CreateStrategy(s *Server) gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		log.Info().Msg("handler create strategy")

		var strategy entity.Strategy
		err := c.BindJSON(&strategy)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		err = s.UsecaseService.CreateStrategy(&strategy)
		if err != nil {
			log.Info().Msgf("error on list strategies %v", err.Error())
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, strategy)
	})
}
