package webRouter

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ginHands struct {
	*slog.Logger
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
}

func Logger(serName string, l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// after request
		// latency := time.Since(t)
		// clientIP := c.ClientIP()
		// method := c.Request.Method
		// statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		cData := &ginHands{
			Logger:     l,
			SerName:    serName,
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			MsgStr:     msg,
		}

		c.Next()

		logSwitch(cData)
	}
}

func logSwitch(data *ginHands) {
	var e func(msg string, args ...any)

	switch {
	case data.StatusCode >= http.StatusBadRequest && data.StatusCode <= http.StatusInternalServerError:
		e = data.Warn
	default:
		e = data.Info
	}

	e("route handled", "status", data.StatusCode, "method", data.Method, "endpoint", data.Path, "latency", data.Latency)
}
