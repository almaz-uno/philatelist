package imagesearch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	testSearcher struct {
		urls    []string
		err     error
		query   string
		placeid string
	}
	testSearcher1 struct {
		testSearcher
	}
	testSearcher2 struct {
		testSearcher
	}
	testSearcher3 struct {
		testSearcher
	}
	testSearcher4 struct {
		testSearcher
	}
)

var (
	errorEtalon1 = fmt.Errorf("Etalon error1")
	errorEtalon2 = fmt.Errorf("Etalon error2")
)

func (s *testSearcher) SearchByQuery(query string) (urls []string, err error) {
	if s.query != "" {
		panic("reset() invocation is required!")
	}
	s.query = query
	return s.urls, s.err
}
func (s *testSearcher) SearchByPlaceID(placeid string) (urls []string, err error) {
	if s.placeid != "" {
		panic("reset() invocation is required!")
	}
	s.placeid = placeid
	return s.urls, s.err
}

func (s *testSearcher) reset() {
	s.query = ""
	s.placeid = ""
}

func TestSearcher(t *testing.T) {
	assert.Implements(t, (*Searcher)(nil), new(CumulativeSearcher), "imagesearcher.CumulativeSearcher must implements interface!")
}

func TestCumulativeSearcher(t *testing.T) {
	s1 := new(testSearcher1)
	s1.urls = []string{"url1", "url2"}
	s1.err = nil

	s2 := new(testSearcher2)
	s2.urls = []string{"url3", "url4", "url5"}
	s2.err = nil

	s3 := new(testSearcher3)
	s3.urls = nil
	s3.err = errorEtalon1

	s4 := new(testSearcher4)
	s4.urls = nil
	s4.err = errorEtalon2

	resetAll := func() {
		s1.reset()
		s2.reset()
		s3.reset()
		s4.reset()
	}

	t.Run("query-success", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s2)

		urls, err := subj.SearchByQuery("query")
		assert.NoError(t, err)

		assert.Len(t, urls, 5)
		assert.Contains(t, urls, "url1", "url2", "url3", "url4", "url5")

		assert.Equal(t, s1.query, "query")
		assert.Equal(t, s2.query, "query")

	})
	t.Run("query-failed", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s3)
		subj.Add(s4)

		urls, err := subj.SearchByQuery("query")
		assert.Error(t, err)
		assert.Len(t, err, 2)
		assert.Equal(t, "Etalon error1, Etalon error2", err.Error())
		assert.Len(t, urls, 0)

		assert.Equal(t, s3.query, "query")
		assert.Equal(t, s4.query, "query")

	})
	t.Run("query-mash-solitary", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s3)

		urls, err := subj.SearchByQuery("query")
		assert.Error(t, err)
		assert.Equal(t, errorEtalon1, err)
		assert.Len(t, urls, 2)
		assert.Contains(t, urls, "url1", "url2")

		assert.Equal(t, s1.query, "query")
		assert.Equal(t, s3.query, "query")

	})
	t.Run("query-mash-fullhouse", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s2)
		subj.Add(s3)
		subj.Add(s4)
		subj.Add(s3) // <= premeditated
		subj.Add(s4) // <= premeditated

		require.Len(t, subj.searchers, 4)

		urls, err := subj.SearchByQuery("query")

		assert.Error(t, err)
		assert.Len(t, err, 2)
		assert.Contains(t, err, errorEtalon1, errorEtalon2)
		assert.Len(t, urls, 5)
		assert.Contains(t, urls, "url1", "url2", "url3", "url4", "url5")

		assert.Equal(t, s1.query, "query")
		assert.Equal(t, s2.query, "query")
		assert.Equal(t, s3.query, "query")
		assert.Equal(t, s4.query, "query")

	})

	t.Run("placeid-success", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s2)

		urls, err := subj.SearchByPlaceID("placeid")
		assert.NoError(t, err)

		assert.Len(t, urls, 5)
		assert.Contains(t, urls, "url1", "url2", "url3", "url4", "url5")

		assert.Equal(t, s1.placeid, "placeid")
		assert.Equal(t, s2.placeid, "placeid")

	})
	t.Run("placeid-failed", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s3)
		subj.Add(s4)

		urls, err := subj.SearchByPlaceID("placeid")
		assert.Error(t, err)
		assert.Len(t, err, 2)
		assert.Equal(t, "Etalon error1, Etalon error2", err.Error())
		assert.Len(t, urls, 0)

		assert.Equal(t, s3.placeid, "placeid")
		assert.Equal(t, s4.placeid, "placeid")

	})
	t.Run("placeid-mash-solitary", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s3)

		urls, err := subj.SearchByPlaceID("placeid")
		assert.Error(t, err)
		assert.Equal(t, errorEtalon1, err)
		assert.Len(t, urls, 2)
		assert.Contains(t, urls, "url1", "url2")

		assert.Equal(t, s1.placeid, "placeid")
		assert.Equal(t, s3.placeid, "placeid")

	})
	t.Run("placeid-mash-fullhouse", func(t *testing.T) {
		resetAll()
		subj := new(CumulativeSearcher)

		subj.Add(s1)
		subj.Add(s2)
		subj.Add(s3)
		subj.Add(s4)
		subj.Add(s3) // <= premeditated
		subj.Add(s4) // <= premeditated

		require.Len(t, subj.searchers, 4)

		urls, err := subj.SearchByPlaceID("placeid")

		assert.Error(t, err)
		assert.Len(t, err, 2)
		assert.Contains(t, err, errorEtalon1, errorEtalon2)
		assert.Len(t, urls, 5)
		assert.Contains(t, urls, "url1", "url2", "url3", "url4", "url5")

		assert.Equal(t, s1.placeid, "placeid")
		assert.Equal(t, s2.placeid, "placeid")
		assert.Equal(t, s3.placeid, "placeid")
		assert.Equal(t, s4.placeid, "placeid")

	})

}
