// Copyright © 2017 Maxim Kovrov <maksim.kovrov@gmail.com>
//

package cmd

import (
	"fmt"
	"net/http/httputil"
	"strings"
	"time"

	"bitbucket.org/CuredPlumbum/philatelist/google"
	"bitbucket.org/CuredPlumbum/philatelist/imagesearch"

	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs RESTful server",
	Long:  `This command launches RESTful server for searching images by location.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Running...")
		err := run()
		if err != nil {
			log.WithError(err).Error("Error while executing run")
			return
		}

	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	thisCmd := runCmd

	thisCmd.PersistentFlags().StringP("listen", "l", "192.168.2.50:9080", "Interface and port for expose RESTful interface")
	thisCmd.PersistentFlags().StringP("google-key", "G", "", "Google API key. To gain new one, please visit https://support.google.com/googleapi/answer/6158862. Required.")
	thisCmd.PersistentFlags().StringP("google-timeout", "T", "30s", "Timeout to Google API invocation. Value must be parsable with `time.ParseDuration` function.")

	mustBindPFlag("run.listen", thisCmd.PersistentFlags().Lookup("listen"))
	mustBindPFlag("run.google.key", thisCmd.PersistentFlags().Lookup("google-key"))
	mustBindPFlag("run.google.timeout", thisCmd.PersistentFlags().Lookup("google-timeout"))

}

func run() (err error) {
	listenOn := strings.TrimSpace(viper.GetString("run.listen"))
	googleKey := strings.TrimSpace(viper.GetString("run.google.key"))
	timeoutStr := strings.TrimSpace(viper.GetString("run.google.timeout"))

	if googleKey == "" {
		return fmt.Errorf("please, specify Google API key; see https://support.google.com/googleapi/answer/6158862 for additional info")
	}

	googleTimeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return fmt.Errorf("error while parsing string '%v' as duration. Error is: %v", timeoutStr, err)
	}

	// construct searchers

	cumulativeSearcher := &imagesearch.CumulativeSearcher{}

	gs := google.New(googleKey)
	gs.Timeout = googleTimeout
	cumulativeSearcher.Add(gs)

	service := imagesearch.NewRestService(cumulativeSearcher)

	// exit hook
	go exitHook()

	// run
	e := echo.New()
	addHTTPDumpMiddleware(e)

	service.RegisterOperations(e)

	err = e.Start(listenOn)
	if err != nil {
		return
	}

	return
}

func addHTTPDumpMiddleware(e *echo.Echo) {
	if log.GetLevel() >= log.DebugLevel {
		e.Use(httpDumpMiddleware)
	}
}

func httpDumpMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		// dump the request
		if dump, err := httputil.DumpRequest(c.Request(), true); err == nil {
			log.Debug("incoming req:", string(dump))
		} else {
			log.WithError(err).Warn("Error occurred while dumping http request")
		}

		if err := next(c); err != nil {
			c.Error(err)
		}

		log.Debug("Response size:", c.Response().Size)

		return nil
	}
}
