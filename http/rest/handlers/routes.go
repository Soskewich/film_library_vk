package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Register(r *mux.Router, lg *logrus.Logger, db *sqlx.DB) {
	handler := newHandler(lg, db)
	r.HandleFunc("/actor", handler.Create()).Methods(http.MethodPost)
	r.HandleFunc("/actor/{id}", handler.Get()).Methods(http.MethodGet)
	r.HandleFunc("/actor/{id}", handler.Update()).Methods(http.MethodPut)
	r.HandleFunc("/actor/{id}", handler.Delete()).Methods(http.MethodDelete)
}
