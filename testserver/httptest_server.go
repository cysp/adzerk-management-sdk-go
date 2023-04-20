package testserver

import (
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHttpTestServer() *httptest.Server {
	r := chi.NewRouter()

	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	addAdTypeRouteHandlers(r)
	addChannelRouteHandlers(r)
	addSiteRouteHandlers(r)
	addChannelSiteMapRouteHandlers(r)

	return httptest.NewServer(r)
}
