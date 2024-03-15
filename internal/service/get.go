package service

import (
	"context"
	"errors"
	"github.com/Soskewich/film_library_vk/internal/model"
	"github.com/Soskewich/film_library_vk/pkg/db"
	"github.com/Soskewich/film_library_vk/pkg/erru"
)

func (s Service) Get(ctx context.Context, id int) (model.Actor, error) {
	todo, err := s.repo.FindActor(ctx, id)
	switch {
	case err == nil:
	case errors.As(err, &db.ErrObjectNotFound{}):
		return model.Actor{}, erru.ErrArgument{errors.New("actor object not found")}
	default:
		return model.Actor{}, err
	}
	return todo, nil
}
