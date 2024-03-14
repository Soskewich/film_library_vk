package model

import "time"

type Film struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	DateExit    time.Time `json:"dateExit"`
	Rating      int       `json:"rating"`
}

type FilmActor struct {
	Id    int    `json:"id"`
	Actor *Actor `json:"actor,omitemoty"`
	Film  *Film  `json:"film,omitempty"`
}
