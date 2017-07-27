// Package imagesearch contains common structs, interface and other useful stuff for image searching process
package imagesearch

import "strings"

type (
	// Searcher represents API for image searching
	Searcher interface {
		// SearchByQuery returns images URLs connected with palace described by the `query`
		SearchByQuery(query string) (urls []string, err error)
		// SearchByPlaceID returns images URLs connected with palace described by the `placeid`
		SearchByPlaceID(placeid string) (urls []string, err error)
	}

	// CumulativeSearcher implements `Searcher` interface with using set of searchers
	CumulativeSearcher struct {
		searchers []Searcher
	}

	// Troubles just represents a number of errors
	Troubles []error
)

func (errs Troubles) Error() string {
	var strRep = make([]string, len(errs))
	for i, e := range errs {
		strRep[i] = e.Error()
	}
	return strings.Join(strRep, ", ")
}

// AsError casts Troubles to an error or nil, if there are no errors
func (errs Troubles) AsError() error {
	switch {
	case len(errs) == 1:
		return errs[0]
	case len(errs) > 1:
		return errs
	}
	return nil
}

// Add allows to add a searcher to the searchers pool
// The method returns this CumulativeSearcher for chain invocation if needed
func (s *CumulativeSearcher) Add(searcher Searcher) {
	for _, es := range s.searchers {
		if es == searcher {
			return
		}
	}
	s.searchers = append(s.searchers, searcher)
}

// SearchByQuery searches images with searchers those added with `Add` method
// Method can returns errors and non-nil urls simultaneously. In that case urls can be safety used,
// but error can be printed to log, for example
// `urls` should be tested against `len(urls) > 0` for output to client
func (s *CumulativeSearcher) SearchByQuery(query string) (urls []string, err error) {
	var errs Troubles
	for _, serch := range s.searchers {
		u, e := serch.SearchByQuery(query)
		urls = append(urls, u...)
		if e != nil {
			errs = append(errs, e)
		}
	}

	err = errs.AsError()

	return
}

// SearchByPlaceID ищет изображения с помощью заданных searchers
// Method can returns errors and non-nil urls simultaneously. In that case urls can be safety used,
// but error can be printed to log, for example
// `urls` should be tested against `len(urls) > 0` for output to client
func (s *CumulativeSearcher) SearchByPlaceID(placeid string) (urls []string, err error) {
	var errs Troubles
	for _, serch := range s.searchers {
		u, e := serch.SearchByPlaceID(placeid)
		urls = append(urls, u...)
		if e != nil {
			errs = append(errs, e)
		}
	}

	err = errs.AsError()

	return
}
