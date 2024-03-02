package webRouter

import (
	"time"

	"github.com/gin-contrib/cors"
)

func (s *Service) initCorsMiddleware() {
	if s.cors == nil {
		// Configure CORS middleware options
		config := cors.Config{
			AllowOrigins:     []string{"*"}, // Vue app's address, adjust as needed
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost:8080" // Use this to allow specific origins
			},
			MaxAge: 12 * time.Hour,
		}

		s.cors = cors.New(config)
	}

	s.api.Use(s.cors)
}
