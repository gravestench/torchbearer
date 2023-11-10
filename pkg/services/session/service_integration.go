package session

type SessionManager interface {
	LoadSessions() error
	SaveSessions() error
	Sessions() (map[string]Session, error)
	NewSession(name string) Session
	BeginSession(name string) error
	EndSession()
}
