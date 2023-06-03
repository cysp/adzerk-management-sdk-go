package testserver

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	adzerk "github.com/cysp/adzerk-management-sdk-go"
)

func addUserRouteHandlers(r *chi.Mux) {
	users := make(map[int64]*adzerk.User)

	var userIdCounter = int64(900_000)
	newUserId := func() int64 {
		userIdCounter = userIdCounter + 1
		return userIdCounter
	}

	r.Post("/v1/login", func(w http.ResponseWriter, r *http.Request) {
		var rb adzerk.CreateUserJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userId := newUserId()
		users[userId] = &adzerk.User{Id: int32(userId)}

		writeJsonMarshalable(w, applyUserCreate(users[userId], rb))
	})

	r.Get("/v1/login/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, users[userId])
	})

	r.Put("/v1/login/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var rb adzerk.UpdateUserJSONRequestBody
		if err := decodeJsonRequestBody(r, &rb); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		writeJsonMarshalable(w, applyUserUpdate(users[userId], rb))
	})

	r.Get("/v1/login/{id}/delete", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		userId, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, found := users[userId]
		if !found {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		delete(users, userId)

		_, err = w.Write([]byte{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func applyUserCreate(user *adzerk.User, update adzerk.CreateUserJSONRequestBody) *adzerk.User {
	user.Email = update.Email
	if update.Name != nil {
		user.Name = *update.Name
	} else {
		user.Name = "null"
	}
	if update.AccessLevel != nil {
		user.AccessLevel = *update.AccessLevel
	} else {
		user.AccessLevel = adzerk.Read
	}
	// user.CanAccessStudio = *update.CanAccessStudio
	return user
}

func applyUserUpdate(user *adzerk.User, update adzerk.UpdateUserJSONRequestBody) *adzerk.User {
	user.Id = update.Id
	user.Email = update.Email
	user.Name = update.Name
	if update.AccessLevel != nil {
		user.AccessLevel = *update.AccessLevel
	} else {
		user.AccessLevel = adzerk.Read
	}
	user.CanAccessStudio = *update.CanAccessStudio
	return user
}
