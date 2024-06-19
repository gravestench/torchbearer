package webRouter

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *Service) initStaticAssetsMiddleware() {
	s.log.Info("setting up routes for static assets")
	if s.api == nil {
		panic("router is nil")
	}

	//s.RouteRoot().Use(gzip.Gzip(gzip.DefaultCompression))
	s.root.NoRoute(s.staticWebUIHandler)
}

func (s *Service) staticWebUIHandler(c *gin.Context) {
	file := c.Request.RequestURI

	_, baseFilename := path.Split(file)
	ext := filepath.Ext(baseFilename)

	//if file == "" || ext == "" {
	//	file = "index.html"
	//}

	data, readErr := s.ReadFile(file)
	if readErr != nil {
		s.log.Warn("reading file", "error", readErr)
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	contentTypeToUse := getMIMEFromFileExtension(ext)
	if contentTypeToUse == "" {
		contentTypeToUse = http.DetectContentType(data)
	}

	c.Data(http.StatusOK, contentTypeToUse, data)

	return
}

func getMIMEFromFileExtension(ext string) (result string) {
	return map[string]string{
		".js":  "text/javascript",
		".css": "text/css",
	}[ext]
}

//go:embed static
var embedded embed.FS

func reformat(p string) string {
	const prefix = "static"

	for {
		if len(p) < 2 {
			break
		}

		if string(p[0]) == "/" {
			p = p[1:]
			continue
		}

		break
	}

	if !strings.HasPrefix(p, "api/") {
		p = fmt.Sprintf("%s/%s", prefix, p)
	}

	return p
}

func (s *Service) Open(name string) (fs.File, error) {
	f, err := embedded.Open(reformat(name))
	if err != nil {
		err = fmt.Errorf("opening file in embedded filesystem: %v", err)
	}

	return f, err
}

func (s *Service) ReadDir(name string) ([]fs.DirEntry, error) {
	list, err := embedded.ReadDir(reformat(name))
	if err != nil {
		err = fmt.Errorf("reading directory in embedded filesystem: %v", err)
	}

	return list, err
}

func (s *Service) ReadFile(name string) ([]byte, error) {
	data, err := embedded.ReadFile(reformat(name))
	if err != nil {
		err = fmt.Errorf("reading file from embedded filesystem: %v", err)
	}

	return data, err
}
