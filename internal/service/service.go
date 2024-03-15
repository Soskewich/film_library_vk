package service

import (
	"github.com/Soskewich/film_library_vk/internal/repository"
)

type Service struct {
	repo repository.ActorRepository
}

func NewService(r repository.ActorRepository) Service {
	return Service{
		repo: r,
	}
}
