package imagesearch

import (
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type (
	// RestService serves RESTful requests for image
	RestService struct {
		rootSearcher Searcher
		basePath     string
	}

	// Image describes an image
	Image struct {
		// URL point to the image
		URL string `json:"url"`
	}
)

// from the Swagger spec
const (
	basePath          = "/v1"
	addressTextPath   = "/images/address-text"
	googlePlaceIDPath = "/images/google-place-id"
)

const (
	errorStringQueryRequired   = "'query' parameter is required"
	errorStringPlaceidRequired = "'placeid' parameter is required"
	errorStringImagesNotFound  = "images not found"
)

// NewRestService creates new service
func NewRestService(rootSearcher Searcher) *RestService {
	return &RestService{
		rootSearcher: rootSearcher,
		basePath:     basePath,
	}
}

// RegisterOperations registers operations in echo instance
func (rs *RestService) RegisterOperations(e *echo.Echo) {
	log.Info("Register operations...")
	log.Info("Registring ", basePath+addressTextPath)
	e.GET(rs.basePath+addressTextPath, rs.doAddressText)
	log.Info("Registring ", basePath+googlePlaceIDPath)
	e.GET(rs.basePath+googlePlaceIDPath, rs.doGooglePlaceID)
}

func (rs *RestService) doAddressText(cxt echo.Context) (err error) {
	query := cxt.QueryParam("query")
	if query == "" {
		log.WithField("url", cxt.Request().URL).Warn("Required parameter 'query' is not specified")
		return echo.NewHTTPError(http.StatusBadRequest, errorStringQueryRequired)
	}

	urls, err := rs.rootSearcher.SearchByQuery(query)
	if err != nil {
		log.WithError(err).WithField("url", cxt.Request().URL).Error("Error while executing request")
	}

	if len(urls) == 0 {
		log.WithField("url", cxt.Request().URL).Warn("Images not found")
		return echo.NewHTTPError(http.StatusNotFound, errorStringImagesNotFound)
	}

	res := make([]*Image, len(urls))
	for i, u := range urls {
		res[i] = &Image{URL: u}
	}
	return cxt.JSONPretty(http.StatusOK, res, "  ")
}
func (rs *RestService) doGooglePlaceID(cxt echo.Context) (err error) {
	placeid := cxt.QueryParam("placeid")
	if placeid == "" {
		log.WithField("url", cxt.Request().URL).Warn("Required parameter 'placeid' is not specified")
		return echo.NewHTTPError(http.StatusBadRequest, errorStringPlaceidRequired)
	}

	urls, err := rs.rootSearcher.SearchByPlaceID(placeid)
	if err != nil {
		log.WithError(err).WithField("url", cxt.Request().URL).Error("Error while executing request")
	}

	if len(urls) == 0 {
		log.WithField("url", cxt.Request().URL).Warn("Images not found")
		return echo.NewHTTPError(http.StatusNotFound, errorStringImagesNotFound)
	}

	res := make([]*Image, len(urls))
	for i, u := range urls {
		res[i] = &Image{URL: u}
	}
	return cxt.JSONPretty(http.StatusOK, res, "  ")
}
