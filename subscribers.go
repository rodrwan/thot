package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Subscriber ...
type Subscriber struct {
	Endpoint string
	Method   string
	Name     string
	URL      string
	handler  http.HandlerFunc
}

// Subscribers ...
type Subscribers map[string]*Subscriber

func (s Subscribers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if _, ok := s[path]; ok {
		s[path].handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

// HandleFunc ...
func (s Subscribers) HandleFunc(mux *mux.Router, subs *Subscriber, handler http.HandlerFunc) {
	key := fmt.Sprintf("/%s/%s", strings.ToLower(subs.Name), subs.Endpoint)
	s[key] = subs
	s[key].handler = handler
	mux.HandleFunc(key, s.ServeHTTP).
		Methods(subs.Method).
		Name(subs.Name)
}
