package testserver

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	adzerk "github.com/cysp/adzerk-management-sdk-go"
)

func addChannelRouteHandlers(r *chi.Mux) {
	channels := make(map[int64]*adzerk.Channel)

	var channelIdCounter = int64(100_000)
	newChannelId := func() int64 {
		channelIdCounter = channelIdCounter + 1
		return channelIdCounter
	}

	r.Post("/v1/channel", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.CreateChannelJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		channelId := newChannelId()
		engine := "0"
		channels[channelId] = &adzerk.Channel{Id: int32(channelId), Engine: &engine, AdTypes: []int32{}}

		writeJsonMarshalable(w, applyChannelCreate(channels[channelId], rb))
	})

	r.Get("/v1/channel/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		channelId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, channels[channelId])
	})

	r.Put("/v1/channel/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		channelId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var rb adzerk.UpdateChannelJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, applyChannelUpdate(channels[channelId], rb))
	})

	r.Get("/v1/channel/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		channelId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, found := channels[channelId]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		delete(channels, channelId)

		_, err = w.Write([]byte{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})
}

func applyChannelCreate(channel *adzerk.Channel, create adzerk.CreateChannelJSONRequestBody) *adzerk.Channel {
	if channel == nil {
		return nil
	}
	channel.Title = create.Title
	channel.AdTypes = create.AdTypes
	return channel
}

func applyChannelUpdate(channel *adzerk.Channel, update adzerk.UpdateChannelJSONRequestBody) *adzerk.Channel {
	if channel == nil {
		return nil
	}
	channel.Id = update.Id
	channel.Title = update.Title
	channel.AdTypes = update.AdTypes
	return channel
}
