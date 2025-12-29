package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rafaeldepontes/gopher-sub/internal/auth"
	"github.com/rafaeldepontes/gopher-sub/internal/auth/model"
	"github.com/rafaeldepontes/gopher-sub/internal/auth/service"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/sirupsen/logrus"
)

type authController struct {
	authService auth.Service
	log         *logrus.Logger
}

func NewController() auth.Controller {
	return &authController{
		authService: service.NewService(),
		log:         logger.GetLogger(),
	}
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var userReq model.UserReq
	json.NewDecoder(r.Body).Decode(&userReq)

	if err := c.authService.Login(&userReq); err != nil {
		c.log.Errorln("Something went wrong log in: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Do all the token logic after...
}

func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	var userReq model.UserReq
	json.NewDecoder(r.Body).Decode(&userReq)

	if err := c.authService.Register(&userReq); err != nil {
		c.log.Errorln("Something went wrong on registering: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
