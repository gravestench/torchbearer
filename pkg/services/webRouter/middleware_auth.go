package webRouter

func (s *Service) initAuthMiddleware() {
	for _, service := range s.mesh.Services() {
		if candidate, ok := service.(ProvidesAuthMiddleware); ok {
			s.auth = append(s.auth, candidate.AuthMiddleware())
		}
	}
}
