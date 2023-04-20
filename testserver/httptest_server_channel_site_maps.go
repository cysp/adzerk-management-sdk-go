package testserver

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	adzerk "github.com/cysp/adzerk-management-sdk-go"
)

type channelSiteMapKey struct {
	channelId int32
	siteId    int32
}

func addChannelSiteMapRouteHandlers(r *chi.Mux) {
	channelSiteMapsByKey := make(map[channelSiteMapKey]*adzerk.ChannelSiteMap)

	r.Post("/v1/channelSite", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.CreateChannelSiteMapJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		key := channelSiteMapKey{channelId: rb.ChannelId, siteId: rb.SiteId}

		existing, found := channelSiteMapsByKey[key]
		if found {
			writeJsonMarshalable(w, applyChannelSiteMapCreate(existing, rb))
			return
		}

		channelSiteMap := adzerk.ChannelSiteMap{ChannelId: rb.ChannelId, SiteId: rb.SiteId, Priority: &rb.Priority}

		channelSiteMapsByKey[key] = &channelSiteMap

		writeJsonMarshalable(w, applyChannelSiteMapCreate(&channelSiteMap, rb))
	})

	r.Get("/v1/channel/{channelId}/site/{siteId}", func(w http.ResponseWriter, r *http.Request) {
		channelId, err := strconv.ParseInt(chi.URLParam(r, "channelId"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		siteId, err := strconv.ParseInt(chi.URLParam(r, "siteId"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		key := channelSiteMapKey{channelId: int32(channelId), siteId: int32(siteId)}

		existing, found := channelSiteMapsByKey[key]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		writeJsonMarshalable(w, existing)

	})

	r.Put("/v1/channelSite", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.UpdateChannelSiteMapJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		key := channelSiteMapKey{channelId: rb.ChannelId, siteId: rb.SiteId}

		existing, found := channelSiteMapsByKey[key]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		writeJsonMarshalable(w, applyChannelSiteMapUpdate(existing, rb))
	})

	r.Get("/v1/channel/{channelId}/site/{siteId}/delete", func(w http.ResponseWriter, r *http.Request) {
		channelId, err := strconv.ParseInt(chi.URLParam(r, "channelId"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		siteId, err := strconv.ParseInt(chi.URLParam(r, "siteId"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		key := channelSiteMapKey{channelId: int32(channelId), siteId: int32(siteId)}

		channelSiteMap, found := channelSiteMapsByKey[key]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		delete(channelSiteMapsByKey, channelSiteMapKey{channelId: channelSiteMap.ChannelId, siteId: channelSiteMap.SiteId})

		_, err = w.Write([]byte{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func applyChannelSiteMapCreate(channelSiteMap *adzerk.ChannelSiteMap, update adzerk.CreateChannelSiteMapJSONRequestBody) *adzerk.ChannelSiteMap {
	channelSiteMap.ChannelId = update.ChannelId
	channelSiteMap.SiteId = update.SiteId
	channelSiteMap.Priority = &update.Priority
	return channelSiteMap
}

func applyChannelSiteMapUpdate(channelSiteMap *adzerk.ChannelSiteMap, update adzerk.UpdateChannelSiteMapJSONRequestBody) *adzerk.ChannelSiteMap {
	channelSiteMap.ChannelId = update.ChannelId
	channelSiteMap.SiteId = update.SiteId
	channelSiteMap.Priority = &update.Priority
	return channelSiteMap
}
