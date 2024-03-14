package repository

import (
	"context"
	"fmt"
	"github.com/Soskewich/film_library_vk/db"
	"github.com/Soskewich/film_library_vk/internal/model"

	"github.com/jmoiron/sqlx"
)

type ActorRepository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) ActorRepository {
	return ActorRepository{Db: db}
}

func (ar ActorRepository) FindActor(ctx context.Context, id int) (model.Actor, error) {
	entity := model.Actor{}
	query := fmt.Sprintf(
		"SELECT * FROM actor WHERE id = $1 IS NULL",
	)
	err := ar.Db.GetContext(ctx, &entity, query, id)
	return entity, db.HandleError(err)
}

func (ar ActorRepository) CreateActor(ctx context.Context, entity *model.Actor) error {
	query := `INSERT INTO actor (name, surname, patronymic, birthday, gender)
                VALUES (:name, :surname, :patronymic, :birthday, :gender) RETURNING id;`
	rows, err := ar.Db.NamedQueryContext(ctx, query, entity)
	if err != nil {
		return db.HandleError(err)
	}

	for rows.Next() {
		err = rows.StructScan(entity)
		if err != nil {
			return db.HandleError(err)
		}
	}
	return db.HandleError(err)
}

func (ar ActorRepository) Update(ctx context.Context, entity model.Actor) error {
	query := `UPDATE actor
                SET name = :name, 
                    surname = :surname, 
                    patronymic = :patronymic, 
                    birthday = :birthday, 
                    gender = :gender
                WHERE id = :id;`
	_, err := ar.Db.NamedExecContext(ctx, query, entity)
	return db.HandleError(err)
}
