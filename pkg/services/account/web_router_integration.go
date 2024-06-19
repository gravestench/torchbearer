package account

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "account"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.POST("login", s.handleLogin)
	group.POST("logout", s.handleLogout)
	group.GET("authenticated", s.handleCheckAuthenticated)
	group.POST("create", s.handleCreate)
	group.POST("recover", s.handleSendOneTimePassCode)
}

func (s *Service) handleCreate(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user name required"})
		return
	}

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "email required"})
		return
	}

	if _, err := s.CreateAccount(username, email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account created"})
}

func (s *Service) handleLogin(c *gin.Context) {
	type payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	data := &payload{}
	c.Bind(data)

	account, err := s.GetAccountByName(data.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("login failed: %v", err)})
		return
	}

	// Validate username and password (you should use a database here)
	if account.Username == data.Username && data.Password == account.Password {
		session := sessions.Default(c)
		session.Set("authenticated", true)
		session.Set("uuid", account.ID.String())
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
	}
}

func (s *Service) handleLogout(c *gin.Context) {
	type payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	data := &payload{}
	c.Bind(data)

	account, err := s.GetAccountByName(data.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("login failed: %v", err)})
		return
	}

	// Validate username and password (you should use a database here)
	if account.Username == data.Username && data.Password == account.Password {
		session := sessions.Default(c)
		session.Set("authenticated", false)
		session.Delete("uuid")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Logout failed"})
	}
}

func (s *Service) handleCheckAuthenticated(c *gin.Context) {
	const sessionTimeout = time.Second * 5

	session := sessions.Default(c)

	if s.accountTimeout == nil {
		s.accountTimeout = make(map[string]time.Time)
	}

	if session.Get("uuid") == nil {
		c.Status(http.StatusUnauthorized)
		return
	}

	id := session.Get("uuid").(string)

	if id == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	if lastTime, found := s.accountTimeout[id]; found {
		if time.Since(lastTime) > sessionTimeout {
			session.Delete("authenticated")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Session timed out"})
			return
		}
	}

	s.accountTimeout[id] = time.Now()

	auth := session.Get("authenticated")
	switch v := auth.(type) {
	case bool:
		if v {
			c.Status(http.StatusOK)
			return
		}
	}

	c.Status(http.StatusUnauthorized)
}

func (s *Service) handleSendOneTimePassCode(c *gin.Context) {
	email := c.PostForm("email")

	if email == "" {
		c.String(http.StatusBadRequest, fmt.Sprintf("malformed payload: invalid email"))
		return
	}

	if err := s.SendOneTimePassCode(email); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("sending one-time pass code: %v", err))
		return
	}

	c.Status(http.StatusOK)
}
