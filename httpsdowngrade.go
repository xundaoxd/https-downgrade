// Package httpsdowngrade a HTTPDowngrade plugin.
package httpsdowngrade

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
    Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
        Headers: make(map[string]string),
    }
}

// HTTPDowngrade a HTTPDowngrade plugin.
type HTTPDowngrade struct {
	next     http.Handler
	name     string
}

// New created a new HTTPDowngrade plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HTTPDowngrade{
		next:     next,
		name:     name,
	}, nil
}

func (a *HTTPDowngrade) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	a.next.ServeHTTP(rw, req)
}
