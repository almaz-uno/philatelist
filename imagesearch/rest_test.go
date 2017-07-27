package imagesearch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	tu "github.com/cured-plumbum/testutil"
	"github.com/labstack/echo"
	"github.com/smallfish/simpleyaml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	tu.InitLog()
}

func TestRegisterOperations(t *testing.T) {

	ts := new(testSearcher)
	ts.urls = []string{"url1", "url2", "url3"}

	rs := NewRestService(ts)
	e := echo.New()
	rs.RegisterOperations(e)

	// Test against swagger specification
	swaggerSpecFile := tu.EvalPath("/iqhive/src/bitbucket.org/CuredPlumbum/philatelist/imagesearch",
		"/iqhive/src/bitbucket.org/CuredPlumbum/philatelist/philatelist.swagger.yaml")

	specBytes, err := ioutil.ReadFile(swaggerSpecFile)
	require.NoError(t, err)

	yaml, err := simpleyaml.NewYaml(specBytes)
	require.NoError(t, err)

	paths, err := yaml.Get("paths").GetMapKeys()
	require.NoError(t, err)

	t.Log("paths in spec:")
	for _, path := range paths {
		t.Log("\t", path)
	}

	var routes = make(map[string]bool)

	// folging by path
	for _, v := range e.Routes() {
		routes[v.Path] = true
	}

	assert.Len(t, reflect.ValueOf(routes).MapKeys(), len(paths), "Number of routers doesn't correspond number of paths. See 'paths' section in the spec.")

	for _, route := range e.Routes() {
		path := ""
		found := false
		for _, p := range paths {
			if rs.basePath+p == route.Path {
				path = p
				found = true
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("Path for route %v not found in the spec", route.Path))
		_, err := yaml.GetPath("paths", path, strings.ToLower(route.Method)).GetMapKeys()
		assert.NoError(t, err, "Can not found in swagger spec "+route.Method+" method for route "+route.Path)
	}

}

func TestOperations(t *testing.T) {

	const successfullyReply = `[
  {
    "url": "url1"
  },
  {
    "url": "url2"
  },
  {
    "url": "url3"
  }
]`

	ts := new(testSearcher)
	ts.urls = []string{"url1", "url2", "url3"}

	rs := NewRestService(ts)
	e := echo.New()

	t.Run("address-text-successfully", func(t *testing.T) {

		req := httptest.NewRequest(echo.GET, "/?query=some+free-form+address", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs.doAddressText(cxt)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		resBody := string(rec.Body.Bytes())

		assert.Equal(t, successfullyReply, resBody, tu.DiffForAssert(successfullyReply, resBody))

	})
	t.Run("address-text-failed", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs.doAddressText(cxt)
		require.Equal(t, echo.NewHTTPError(http.StatusBadRequest, "'query' parameter is required"), err)
	})
	t.Run("address-text-not-found", func(t *testing.T) {
		rs404 := NewRestService(new(testSearcher))
		req := httptest.NewRequest(echo.GET, "/?query=nf", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs404.doAddressText(cxt)
		require.Equal(t, echo.NewHTTPError(http.StatusNotFound, "images not found"), err)
	})

	//
	t.Run("placeid-successfully", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/?placeid=placeID123456", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs.doGooglePlaceID(cxt)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, rec.Code)

		resBody := string(rec.Body.Bytes())

		assert.Equal(t, successfullyReply, resBody, tu.DiffForAssert(successfullyReply, resBody))

	})
	t.Run("placeid-failed", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs.doGooglePlaceID(cxt)
		require.Equal(t, echo.NewHTTPError(http.StatusBadRequest, "'placeid' parameter is required"), err)

	})
	t.Run("placeid-not-found", func(t *testing.T) {
		rs404 := NewRestService(new(testSearcher))
		req := httptest.NewRequest(echo.GET, "/?placeid=nf", nil)
		rec := httptest.NewRecorder()
		cxt := e.NewContext(req, rec)

		err := rs404.doGooglePlaceID(cxt)
		require.Equal(t, echo.NewHTTPError(http.StatusNotFound, "images not found"), err)

	})
}
