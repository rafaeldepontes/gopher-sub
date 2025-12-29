package controller

import (
	"net/http"
	"strconv"

	msg "github.com/rafaeldepontes/gopher-sub/internal/errors-messages"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription/service"
	"github.com/sirupsen/logrus"
)

type subController struct {
	subService subscription.Service
	log        *logrus.Logger
}

func NewController() subscription.Controller {
	return &subController{
		subService: service.NewService(),
		log:        logger.GetLogger(),
	}
}

func (c *subController) Subscriber(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")
	if strID == "" {
		c.log.Errorln("Something went wrong with the subscription: Empty id")
		http.Error(w, msg.ErrInvalidId.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.log.Errorln("Something went wrong with the subscription: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := c.subService.Subscribe(id); err != nil {
		c.log.Errorln("Something went wrong with the subscription: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
