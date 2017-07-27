// Package google grant access to certain API Google operations
package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"bitbucket.org/CuredPlumbum/philatelist/imagesearch"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	// searching photos
	// https://developers.google.com/places/web-service/photos
	photoServiceURL = "https://maps.googleapis.com/maps/api/place/photo"
	// inspect specified placeid
	// https://developers.google.com/places/web-service/details
	detailsServiceURL = "https://maps.googleapis.com/maps/api/place/details/json"
	// full text search places (and photo reference)
	// https://developers.google.com/places/web-service/search
	textsearchServiceURL = "https://maps.googleapis.com/maps/api/place/textsearch/json"
)

const (
	paramKey      = "key"
	successStatus = "OK"
)

type (
	// API represents certain Google API facade
	API struct {
		apiKey string
		// Timeout is a timeout Google services invocations
		Timeout time.Duration
	}

	// GeoPoint is a just point on the Earth
	GeoPoint struct {
		Lat float32 `json:"lat"`
		Lng float32 `json:"lng"`
	}
	// Photo is a photo descriptor
	Photo struct {
		HTMLAttributions []string `json:"html_attributions"`
		Height           int      `json:"height"`
		Width            int      `json:"width"`
		PhotoReference   string   `json:"photo_reference"`
	}

	// CommonPart just a common part of Google reponses
	CommonPart struct {
		HTMLAttributions []string `json:"html_attributions"`
		Status           string   `json:"status"`
	}

	// a bit of polymorphism
	statuser interface {
		status() string
	}

	// SearchResponse represents resonse from text search service
	SearchResponse struct {
		Results []Result `json:"results"`
		CommonPart
	}

	// Result is a single result from Google services
	Result struct {
		FormattedAddress string `json:"formatted_address"`
		Geometry         struct {
			Location GeoPoint `json:"location"`
			Viewport *struct {
				Northeast *GeoPoint `json:"northeast"`
				Southwest *GeoPoint `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Photos  []Photo `json:"photos"`
		PlaceID string  `json:"place_id"`
	}

	// DetailsResponse represents resonse from details service
	DetailsResponse struct {
		Result Result `json:"result"`
		CommonPart
	}
)

func (c *CommonPart) status() string {
	return c.Status
}

// New creates a new Google API facade
func New(apiKey string) *API {
	return &API{
		apiKey:  apiKey,
		Timeout: 30 * time.Second,
	}
}

// SearchByPlaceID searches images for specified `placeid` with Google API.
// It searches up to 10, please see google documentation for details
// It returns an empty (`len(urls) == 0`) slice if there are no found images for `placeid`
func (g *API) SearchByPlaceID(placeid string) (urls []string, err error) {
	res, err := g.Details(placeid)
	if err != nil {
		return
	}

	for _, photo := range res.Result.Photos {
		urls = append(urls, g.Photo(photo))
	}

	return
}

// SearchByQuery searches images for address denoted by `query` with Google API.
// At the first, this method searches Google placeids for the address,
// then it uses `SearchByPlaceID` for searching photos.
// It returns an empty (`len(urls) == 0`) slice if there are no found images for the address.
func (g *API) SearchByQuery(query string) (urls []string, err error) {
	var errs imagesearch.Troubles
	res, err := g.TextSearch(query)
	if err != nil {
		return
	}

	for _, gr := range res.Results {
		// we are using placeid only
		u, e := g.SearchByPlaceID(gr.PlaceID)
		if e != nil {
			errs = append(errs, e)
		}

		urls = append(urls, u...)

	}

	err = errs.AsError()

	return
}

// Photo returns an URL, that point to the selected image with original maxwidth
func (g *API) Photo(photo Photo) string {
	return fmt.Sprintf("%s?%s=%s&photoreference=%s&maxwidth=%s",
		photoServiceURL, paramKey, url.QueryEscape(g.apiKey),
		url.QueryEscape(photo.PhotoReference), url.QueryEscape(strconv.Itoa(photo.Width)))
}

// TextSearch does text search for places
// https://developers.google.com/places/web-service/search
func (g *API) TextSearch(query string) (res *SearchResponse, err error) {
	location := textsearchServiceURL + "?" + paramKey + "=" + url.QueryEscape(g.apiKey) + "&query=" + url.QueryEscape(query)
	res = new(SearchResponse)
	err = getResponse(location, g.Timeout, res)
	return
}

// Details returns details about place
// https://developers.google.com/places/web-service/details
func (g *API) Details(placeid string) (res *DetailsResponse, err error) {
	location := detailsServiceURL + "?" + paramKey + "=" + url.QueryEscape(g.apiKey) + "&placeid=" + url.QueryEscape(placeid)
	res = new(DetailsResponse)
	err = getResponse(location, g.Timeout, res)
	return
}

func getResponse(location string, timeout time.Duration, response statuser) (err error) {
	data, err := doGet(location, timeout)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, response)
	if err != nil {
		return
	}

	if response.status() != successStatus {
		err = fmt.Errorf("Google returned status %v", response.status())
		return
	}
	return
}

func doGet(location string, timeout time.Duration) (data []byte, err error) {
	if log.GetLevel() >= log.DebugLevel {
		log.Debug("Request Google with URL:", location)
	}
	status, data, err := fasthttp.GetTimeout(nil, location, timeout)

	if err != nil {
		log.WithError(err).WithField("url", location).WithField("status", status).Error("Error occurred while requesting")
		return
	}

	if log.GetLevel() >= log.DebugLevel {
		log.WithField("url", location).WithField("status", status).WithField("data", string(data)).Debug("go data")
	}

	// I think, we will not rely on HTTP-status, but we must check it
	if status < 200 && status >= 300 {
		err = fmt.Errorf("Unexpected status %v %v", status, http.StatusText(status))
		return
	}
	return
}
