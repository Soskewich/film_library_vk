package model

import (
	"time"
)

type Actor struct {
	Id         int       `json:"id"`
	Name       string    `json:"name" binding:"required"`
	Surname    string    `json:"surname" binding:"required"`
	Patronymic string    `json:"patronymic,omitempty"`
	Birthday   time.Time `json:"birthday" binding:"required"`
	Gender     string    `json:"gender" binding:"required"`
}
