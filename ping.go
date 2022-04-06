package main

import (
	"net/http"
	"regexp"
)

var (
	pingRe = regexp.MustCompile(`^\/ping*$`)
)

type pingHandler struct{}

func (h *pingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && pingRe.MatchString(r.URL.Path):
		h.Ping(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *pingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "alive"}`))
}
