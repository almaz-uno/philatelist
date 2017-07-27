package cmd

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func initLog() {

	var writers []io.Writer
	writers = append(writers, os.Stdout)

	logFile := viper.GetString("log.file")

	if logFile != "-" {
		lumberjackWriter := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    viper.GetInt("log.max-size"), // megabytes
			MaxBackups: viper.GetInt("log.max-backups"),
			MaxAge:     viper.GetInt("log.max-age"), // days
		}
		writers = append(writers, lumberjackWriter)
	}

	level, err := log.ParseLevel(viper.GetString("log.level"))

	if err != nil {
		log.Printf("Invalid parse log level %v. Rollback to info. Error is %v\n", viper.GetString("log.level"), err)
		level = log.InfoLevel
	}

	viper.Set("log.level", level.String())

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{})
	//		log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(io.MultiWriter(writers...))

}

func init() {
	cobra.OnInitialize(initLog)
	RootCmd.PersistentFlags().StringP("log-file", "F", "-", "log file; if this value specified as \"-\" - no output.")
	RootCmd.PersistentFlags().IntP("log-max-size", "S", 30, "maximum log file size in megabytes")
	RootCmd.PersistentFlags().IntP("log-max-backups", "B", 30, "maximum number of log files")
	RootCmd.PersistentFlags().IntP("log-max-age", "A", 30, "maximum log files age in days")
	RootCmd.PersistentFlags().StringP("log-level", "L", "info", "log level: panic|fatal|error|warn|info|debug")

	mustBindPFlag("log.file", RootCmd.PersistentFlags().Lookup("log-file"))
	mustBindPFlag("log.max-size", RootCmd.PersistentFlags().Lookup("log-max-size"))
	mustBindPFlag("log.max-backups", RootCmd.PersistentFlags().Lookup("log-max-backups"))
	mustBindPFlag("log.max-age", RootCmd.PersistentFlags().Lookup("log-max-age"))
	mustBindPFlag("log.level", RootCmd.PersistentFlags().Lookup("log-level"))

}
