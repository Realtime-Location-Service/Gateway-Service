package config

import (
	"time"

	"github.com/spf13/viper"
)

// History ...
type History struct {
	BaseURL        string
	RequestTimeout time.Duration
}

var history = &History{}

// HistoryCfg ...
func HistoryCfg() *History {
	return history
}

// LoadHistoryCfg loads history service configuration
func LoadHistoryCfg() {
	history.BaseURL = viper.GetString("history.base_url")
	history.RequestTimeout = viper.GetDuration("history.request_timeout") * time.Second
}
