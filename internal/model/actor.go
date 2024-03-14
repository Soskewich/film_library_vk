package model

import (
	"database/sql"
	"time"
)

type Actor struct {
	Id         int            `json:"id"`
	Name       string         `json:"name" binding:"required"`
	Surname    string         `json:"surname" binding:"required"`
	Patronymic sql.NullString `json:"patronymic,omitempty"`
	Birthday   time.Time      `json:"birthday" binding:"required"`
	Gender     string         `json:"gender" binding:"required"`
}
