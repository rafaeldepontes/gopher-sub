package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/sirupsen/logrus"
)

type handle struct {
	Log *logrus.Logger
}

func NewHandle() *handle {
	return &handle{
		Log: logger.InitLogger(),
	}
}

func (h *handle) ConfigHandler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)

	r.Route("/api/v1", func(r chi.Router) {

	})
}
