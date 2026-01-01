package middleware

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL")},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: false,
		MaxAge:           300, //5 min
	}))
	r.Use(middleware.StripSlashes)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/login", h.authController.Login)
		r.Post("/register", h.authController.Register)

		r.Group(func(r chi.Router) {
			r.Use(AuthenticationFilter)
			r.Post("/subscribe/{id}", h.subController.Subscriber)
		})
	})
}
