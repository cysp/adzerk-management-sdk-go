package client

import (
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHttpTestServer() *httptest.Server {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Post("/v1/channel", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"Id\":123,\"Title\":\"Title\"}"))
	})
	r.Get("/v1/channel/123", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"Id\":123}"))
	})
	r.Put("/v1/channel/123", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"Id\":123,\"Title\":\"Title\"}"))
	})

	return httptest.NewServer(r)
}
