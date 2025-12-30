package service

import (
	"time"

	"github.com/rafaeldepontes/gopher-sub/internal/email"
	"github.com/rafaeldepontes/gopher-sub/internal/email/service"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription/repository"
	"github.com/sirupsen/logrus"
)

type subService struct {
	emailSvc email.Service
	subRepo  subscription.Repository
	log      *logrus.Logger
}

func NewService() subscription.Service {
	return &subService{
		emailSvc: service.NewService(),
		subRepo:  repository.NewRepository(),
		log:      logger.GetLogger(),
	}
}

// Subscribe implements [subscription.Service].
func (s *subService) Subscribe(id int64) error {
	s.log.Infoln("Subscribing the user:", id)
	// fake a payment service.
	time.Sleep(2 * time.Second)

	// subscribe the user
	if err := s.subRepo.SubscribeUser(id); err != nil {
		return err
	}

	user, err := s.subRepo.FindById(id)
	if err != nil {
		return err
	}

	// sends a email
	if err := s.emailSvc.SendMail(user); err != nil {
		return err
	}

	s.log.Infoln("Subscription done successfully")
	return nil
}
