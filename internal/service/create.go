package service

import (
	"context"
	"database/sql"
	"github.com/Soskewich/film_library_vk/internal/model"
	"github.com/Soskewich/film_library_vk/pkg/erru"
	"github.com/asaskevich/govalidator"
	"time"
)

type ActorParams struct {
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Patronymic string    `json:"patronymic"`
	Birthday   time.Time `json:"birthday"`
	Gender     string    `json:"gender"`
}

func (s Service) CreateActor(ctx context.Context, params ActorParams) (int, error) {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return 0, erru.ErrArgument{Wrapped: err}
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	entity := model.Actor{
		Name:    params.Name,
		Surname: params.Surname,
		Patronymic: sql.NullString{
			String: params.Patronymic,
			Valid:  true,
		},
		Birthday: params.Birthday,
		Gender:   params.Gender,
	}
	err = s.repo.CreateActor(ctx, &entity)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	return entity.Id, err
}
