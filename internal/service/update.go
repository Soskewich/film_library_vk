package service

import (
	"context"
	"database/sql"
	"github.com/Soskewich/film_library_vk/pkg/erru"
	"github.com/asaskevich/govalidator"
	"time"
)

type UpdateParams struct {
	Id         int `valid:"required"`
	Name       *string
	Surname    *string
	Patronymic *string
	Birthday   *time.Time
	Gender     *string
}

func (s Service) Update(ctx context.Context, params UpdateParams) error {
	if _, err := govalidator.ValidateStruct(params); err != nil {
		return erru.ErrArgument{Wrapped: err}
	}

	// find todo object
	actor, err := s.Get(ctx, params.Id)
	if err != nil {
		return err
	}

	if params.Name != nil {
		actor.Name = *params.Name
	}
	if params.Surname != nil {
		actor.Surname = *params.Surname
	}
	if params.Patronymic != nil {
		actor.Patronymic = sql.NullString{
			String: *params.Patronymic,
			Valid:  true,
		}
	}
	if params.Birthday != nil {
		actor.Birthday = *params.Birthday
	}
	if params.Gender != nil {
		actor.Gender = *params.Gender
	}

	tx, err := s.repo.Db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	err = s.repo.Update(ctx, actor)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
