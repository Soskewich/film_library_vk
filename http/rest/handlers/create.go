package handlers

import (
	authorService "github.com/Soskewich/film_library_vk/internal/service"
	"net/http"
	"time"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name       string    `json:"name"`
		Surname    string    `json:"surname"`
		Patronymic string    `json:"patronymic"`
		Birthday   time.Time `json:"birthday"`
		Gender     string    `json:"gender"`
	}

	type response struct {
		Id int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		id, err := s.actorService.CreateActor(r.Context(), authorService.ActorParams{
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
