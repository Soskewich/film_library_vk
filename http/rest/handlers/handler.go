package handlers

import (
	actorRepo "github.com/Soskewich/film_library_vk/internal/repository"
	actorService "github.com/Soskewich/film_library_vk/internal/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger       *logrus.Logger
	router       *mux.Router
	actorService actorService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:       lg,
		actorService: actorService.NewService(actorRepo.NewRepository(db)),
	}
}
