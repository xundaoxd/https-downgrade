// Package plugindemo a demo plugin.
package httpsdowngrade

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// plugin.
type HttpDowngrade struct {
	next     http.Handler
	name     string
}

// New created a new HttpDowngrade plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HttpDowngrade{
		next:     next,
		name:     name,
	}, nil
}

func (a *HttpDowngrade) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	a.next.ServeHTTP(rw, req)
}
