package handlers

import (
	"errors"
	actorService "github.com/Soskewich/film_library_vk/internal/service"
	"github.com/Soskewich/film_library_vk/pkg/erru"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name       *string    `json:"name"`
		Surname    *string    `json:"surname"`
		Patronymic *string    `json:"patronymic"`
		Birthday   *time.Time `json:"birthday"`
		Gender     *string    `json:"gender"`
	}

	type response struct {
		Id int `json:"id"`
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

		req := request{}

		err = s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.actorService.Update(r.Context(), actorService.UpdateParams{
			Id:         id,
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
			Birthday:   req.Birthday,
			Gender:     req.Gender,
		})

		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{Id: id}, http.StatusOK)
	}
}
