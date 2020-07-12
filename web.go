package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"upper.io/db.v3"
)

type Web struct {
	Router *chi.Mux
	DB     db.Database
}

func NewWeb(db db.Database) *Web {
	r := chi.NewRouter()
	wb := &Web{Router: r, DB: db}

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", wb.UserRegister)
		// r.Post("/login")
	})

	return wb
}

func (wb *Web) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wb.Router.ServeHTTP(w, r)
}
