package account

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "account"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) {
	group.POST("login", s.handleLogin)
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
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
	}
}

func (s *Service) handleCheckAuthenticated(c *gin.Context) {
	session := sessions.Default(c)

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
