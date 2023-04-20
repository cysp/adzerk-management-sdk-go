package testserver

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	adzerk "github.com/cysp/adzerk-management-sdk-go"
)

type adTypeKey struct {
	width  int32
	height int32
}

func addAdTypeRouteHandlers(r *chi.Mux) {
	adTypes := make(map[int64]*adzerk.AdType)
	adTypesByKey := make(map[adTypeKey]*adzerk.AdType)

	var adTypeIdCounter = int64(400_000)
	newAdTypeId := func() int64 {
		adTypeIdCounter = adTypeIdCounter + 1
		return adTypeIdCounter
	}

	r.Get("/v1/adtypes", func(w http.ResponseWriter, r *http.Request) {
		adTypesList := make([]adzerk.AdType, len(adTypes))
		for _, v := range adTypes {
			adTypesList = append(adTypesList, *v)
		}

		writeJsonMarshalable(w, adzerk.AdTypeList{
			Page:       1,
			PageSize:   1,
			TotalPages: 1,
			TotalItems: int64(len(adTypesList)),
			Items:      adTypesList,
		})
	})

	r.Post("/v1/adtypes", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.CreateAdTypeJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		key := adTypeKey{width: rb.Width, height: rb.Height}

		existing, found := adTypesByKey[key]
		if found {
			writeJsonMarshalable(w, applyAdTypeCreate(existing, rb))
			return
		}

		adTypeId := newAdTypeId()

		adType := adzerk.AdType{Id: int32(adTypeId)}

		adTypes[adTypeId] = &adType
		adTypesByKey[key] = &adType

		writeJsonMarshalable(w, applyAdTypeCreate(&adType, rb))
	})

	r.Get("/v1/adtypes/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		adTypeId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		adType, found := adTypes[adTypeId]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		delete(adTypes, adTypeId)
		delete(adTypesByKey, adTypeKey{width: adType.Width, height: adType.Height})

		_, err = w.Write([]byte{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func applyAdTypeCreate(adType *adzerk.AdType, update adzerk.CreateAdTypeJSONRequestBody) *adzerk.AdType {
	adType.Width = update.Width
	adType.Height = update.Height
	if update.Name != nil {
		adType.Name = update.Name
	}
	return adType
}
