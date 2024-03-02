package account

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) AuthMiddleware() gin.HandlerFunc {
	const sessionKey = "torchbearer"
	store := cookie.NewStore([]byte(uuid.New().String()))
	s.router.RouteRoot().Use(sessions.Sessions(sessionKey, store))

	return s.authMiddleware
}

func (s *Service) authMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	isAuthenticated := session.Get("authenticated")

	if isAuthenticated != nil && isAuthenticated.(bool) {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
	}
}
