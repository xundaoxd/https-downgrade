// Package plugindemo a demo plugin.
package httpsdowngrade

import (
	"context"
	"net/http"
)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type HttpDowngrade struct {
	next     http.Handler
	name     string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &HttpDowngrade{
		next:     next,
		name:     name,
	}, nil
}

func (a *HttpDowngrade) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	a.next.ServeHTTP(rw, req)
}
