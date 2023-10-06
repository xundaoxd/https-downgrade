// Package httpsdowngrade a HTTPSDowngrade plugin.
package httpsdowngrade

import (
    "context"
    "net/http"
    "os"
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

// HTTPSDowngrade a HTTPSDowngrade plugin.
type HTTPSDowngrade struct {
	next     http.Handler
    headers  map[string]string
	name     string
}

// New created a new HTTPSDowngrade plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HTTPSDowngrade{
        headers:  config.Headers,
		next:     next,
		name:     name,
	}, nil
}

func (a *HTTPSDowngrade) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    os.Stdout.WriteString("HTTPSDowngrade")
	a.next.ServeHTTP(rw, req)
}
