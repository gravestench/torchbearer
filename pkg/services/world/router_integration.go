package world

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "world"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.Worlds)
	})
}
