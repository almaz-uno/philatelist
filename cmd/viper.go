package cmd

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// mustBindPFlag just paincs when BindPFlag returs an error
func mustBindPFlag(key string, flag *pflag.Flag) {
	if err := viper.BindPFlag(key, flag); err != nil {
		panic(err)
	}
	return
}
