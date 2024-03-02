package webRouter

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func (s *Service) initSessionMiddleware() {
	if s.sessions == nil {
		// TODO: make these strings come from somewhere else
		store := cookie.NewStore([]byte("secret"))
		s.sessions = sessions.Sessions("torchbearer", store)
	}

	s.api.Use(s.sessions)
}
