package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaeldepontes/gopher-sub/internal/auth"
	authControl "github.com/rafaeldepontes/gopher-sub/internal/auth/controller"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription"
	subControl "github.com/rafaeldepontes/gopher-sub/internal/subscription/controller"
	"github.com/sirupsen/logrus"
)

type handle struct {
	authController auth.Controller
	subController  subscription.Controller
	Log            *logrus.Logger
}

func NewHandle() *handle {
	return &handle{
		authController: authControl.NewController(),
		subController:  subControl.NewController(),
		Log:            logger.GetLogger(),
	}
}

func (h *handle) ConfigHandler(r *chi.Mux) {
	r.Use(middleware.StripSlashes)

	r.Post("/api/v1/login", h.authController.Login)
	r.Post("/api/v1/register", h.authController.Register)

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(AuthenticationFilter)
		// TODO: Create the email service
		// WIP
		r.Post("/subscribe/{id}", h.subController.Subscriber)
	})
}
