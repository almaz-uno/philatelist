package search

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestInvocation(t *testing.T) {
	assert.True(t, false)

}

func TestGoogleApi(t *testing.T) {
	// https://developers.google.com/places/place-id
	const kovrovKey = "AIzaSyBZeYbJ6pMNUy-VdLVxnhBCwYcWxSrZZAE"
	apiKey := kovrovKey

	t.Run("evaluate-placeid", func(t *testing.T) {
		t.Log("Using API is: ", apiKey)
		var address = ""
		c := &fasthttp.HostClient{
			Addr: "localhost:8080",
		}

		targetUrl := url.Parse("")
	})

	t.Run("photo-by-placeid", func(t *testing.T) {
		// ulitsa Krzhizhanovskogo, Moscow
		const krzhizhanovskogoStreet = "ChIJIY8uUsBMtUYRdCGdlBVZxlU"
		// sanatoriya Podmoskovie
		const sanatoriyaPodmoskovie = "ChIJVVVVVdigSkERIUYqip8z7OU"
		// Novotushinskaya ulitsa
		const novotushinskaya = "ChIJhSOnan5HtUYR5dkSIiDaj-U"

	})

	t.Run("photos-by-query", func(t *testing.T) {
		var address string
		var u url.URL

		address = "поселок санатория подмосковье"

		req := new(fasthttp.Request)
		res := new(fasthttp.Response)

		fasthttp.Get(nil, url)

	})

}
