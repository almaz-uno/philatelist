package search

const (
	// searching photos
	// https://developers.google.com/places/web-service/photos
	photosServiceUrl = "https://maps.googleapis.com/maps/api/place/photo"
	// inspect specified placeid
	// https://developers.google.com/places/web-service/details
	detailsServiceUrl = "https://maps.googleapis.com/maps/api/place/details/json"
	// full text search places (and photo reference)
	// https://developers.google.com/places/web-service/search
	textsearchServiceUrl = "https://maps.googleapis.com/maps/api/place/textsearch/json"
)

const (
	paramKeyApi = ""
)

type GoogleApi struct {
	apiKey string
}

func NewGoogleApi(apiKey string) *GoogleApi {
	return &GoogleApi{
		apiKey: apiKey,
	}
}

func (g *GoogleApi) TextSearch(query string) (items interface{}, err error) {
	return nil, nil
}
