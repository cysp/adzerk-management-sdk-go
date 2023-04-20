package testserver

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	adzerk "github.com/cysp/adzerk-management-sdk-go"
)

func addSiteRouteHandlers(r *chi.Mux) {
	sites := make(map[int64]*adzerk.Site)

	var siteIdCounter = int64(200_000)
	newSiteId := func() int64 {
		siteIdCounter = siteIdCounter + 1
		return siteIdCounter
	}

	r.Post("/v1/site", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.CreateSiteJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		siteId := newSiteId()
		sites[siteId] = &adzerk.Site{Id: int32(siteId)}

		writeJsonMarshalable(w, applySiteCreate(sites[siteId], rb))
	})

	r.Get("/v1/site/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		siteId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, sites[siteId])
	})

	r.Put("/v1/site/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		siteId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var rb adzerk.UpdateSiteJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, applySiteUpdate(sites[siteId], rb))
	})

	r.Get("/v1/site/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		siteId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, found := sites[siteId]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		delete(sites, siteId)

		_, err = w.Write([]byte{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func applySiteCreate(site *adzerk.Site, update adzerk.CreateSiteJSONRequestBody) *adzerk.Site {
	site.Title = update.Title
	site.Url = update.URL
	return site
}

func applySiteUpdate(site *adzerk.Site, update adzerk.UpdateSiteJSONRequestBody) *adzerk.Site {
	site.Id = update.Id
	site.Title = update.Title
	site.Url = update.URL
	return site
}
