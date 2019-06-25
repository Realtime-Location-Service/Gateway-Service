package config

import (
	"time"

	"github.com/spf13/viper"
)

// Ping ...
type Ping struct {
	BaseURL        string
	RequestTimeout time.Duration
}

var ping = &Ping{}

// PingCfg ...
func PingCfg() *Ping {
	return ping
}

// LoadPingCfg loads ping configuration
func LoadPingCfg() {
	ping.BaseURL = viper.GetString("ping.base_url")
	ping.RequestTimeout = viper.GetDuration("ping.request_timeout") * time.Second
}
