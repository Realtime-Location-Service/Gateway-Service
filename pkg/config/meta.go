package config

import (
	"time"

	"github.com/spf13/viper"
)

// Meta ...
type Meta struct {
	BaseURL        string
	RequestTimeout time.Duration
}

var meta = &Meta{}

// MetaCfg ...
func MetaCfg() *Meta {
	return meta
}

// LoadMetaCfg loads metadata service configuration
func LoadMetaCfg() {
	meta.BaseURL = viper.GetString("meta.base_url")
	meta.RequestTimeout = viper.GetDuration("meta.request_timeout") * time.Second
}
