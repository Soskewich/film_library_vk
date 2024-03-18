package handlers

import (
	"errors"
	"github.com/Soskewich/film_library_vk/pkg/erru"

	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s service) Get() http.HandlerFunc {
	type response struct {
		Id         int       `json:"id"`
		Name       string    `json:"name"`
		Surname    string    `json:"surname"`
		Patronymic string    `json:"patronymic"`
		Birthday   time.Time `json:"birthday"`
		Gender     string    `json:"gender"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		getResponse, err := s.actorService.Get(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			Id:         getResponse.Id,
			Name:       getResponse.Name,
			Surname:    getResponse.Surname,
			Patronymic: getResponse.Patronymic,
			Birthday:   getResponse.Birthday,
			Gender:     getResponse.Gender,
		}, http.StatusOK)
	}
}
