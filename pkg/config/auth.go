package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Auth ...
type Auth struct {
	RequestTimeout time.Duration
	BaseURL        string
	AuthURL        string
}

var auth = &Auth{}

// AuthCfg ...
func AuthCfg() *Auth {
	return auth
}

// LoadAuthCfg ....
func LoadAuthCfg() {
	auth.BaseURL = viper.GetString("auth.base_url")
	auth.AuthURL = fmt.Sprintf("%s/%s", auth.BaseURL, viper.GetString("auth.authorization"))
	auth.RequestTimeout = viper.GetDuration("auth.request_timeout") * time.Second
}
