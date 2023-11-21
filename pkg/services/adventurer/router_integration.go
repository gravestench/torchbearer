package adventurer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "adventurer"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.GET("procedures", func(c *gin.Context) {
		obj := (&procedureCreateAdventurer{service: s}).New()
		c.JSON(http.StatusOK, obj)
	})
}
