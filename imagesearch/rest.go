package imagesearch

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
)

type (
	// RestService serves RESTful requests for image
	RestService struct {
		rootSearcher Searcher
	}

	// Image describes an image
	Image struct {
		// URL point to the image
		URL string `json:"url"`
	}
)

// NewRestService creates new service
func NewRestService(rootSearcher Searcher) *RestService {
	return &RestService{
		rootSearcher: rootSearcher,
	}
}

// RegisterOperations registers operations in echo instance
func (r *RestService) RegisterOperations(e *echo.Echo) (err error) {
	log.Info("Register operations...")
	panic("NI")
	return
}
