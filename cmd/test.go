// Copyright © 2017 Maxim Kovrov <maksim.kovrov@gmail.com>
//

package cmd

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"bitbucket.org/CuredPlumbum/philatelist/imagesearch"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test the launched RESTful service",
	Long: `This command required for testing of launched service (see command 'run').
Moreover, it demonstrates creation client of the service.
Please warn, service must be running while this command issues.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("test on %v is called", baseURL)

		if placeidRequest == "" && queryRequest == "" {
			log.Error("Please, specify query or placeid")
		}

		if placeidRequest != "" {
			log.Info("Looking up for placeid=", placeidRequest)
			lookup(baseURL + googlePlaceIDPath + "?placeid=" + url.QueryEscape(placeidRequest))
		}
		if queryRequest != "" {
			log.Info("Looking up for query=", queryRequest)
			lookup(baseURL + addressTextPath + "?query=" + url.QueryEscape(queryRequest))
		}

		log.Info("Done!")
	},
}

const (
	googlePlaceIDPath = "/images/google-place-id"
	addressTextPath   = "/images/address-text"
)

var (
	baseURL        string
	queryRequest   string
	placeidRequest string
)

func init() {
	RootCmd.AddCommand(testCmd)
	testCmd.PersistentFlags().StringVarP(&baseURL, "base-url", "b", "http://localhost:9080/v1", "Base target url")

	testCmd.PersistentFlags().StringVarP(&queryRequest, "query", "q", "", "Address query to looking images")
	testCmd.PersistentFlags().StringVarP(&placeidRequest, "placeid", "p", "", "Google placeid to looking images")

}

func lookup(location string) {
	statusCode, body, err := fasthttp.GetTimeout(nil, location, time.Minute)
	if err != nil {
		log.WithError(err).WithField("url", location).Error("Error occurred")
		return
	}

	log.Info("Status code is: ", statusCode)
	log.Info("Body is: ", string(body))

	if statusCode == http.StatusOK {
		var images []imagesearch.Image
		err = json.Unmarshal(body, &images)
		if err != nil {
			log.WithError(err).WithField("url", location).Error("Error while parsing response occurred")
			return
		}
		for _, im := range images {
			log.Info(im.URL)
		}
	}
}
