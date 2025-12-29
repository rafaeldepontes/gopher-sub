package service

import (
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription/repository"
	"github.com/sirupsen/logrus"
)

type subService struct {
	subRepo subscription.Repository
	log     *logrus.Logger
}

func NewService() subscription.Service {
	return &subService{
		subRepo: repository.NewRepository(),
		log:     logger.GetLogger(),
	}
}

// Subscribe implements [subscription.Service].
func (s *subService) Subscribe(id int64) error {
	s.log.Infoln("Subscribing the user:", id)
	// fake a payment service.
	// ...

	// subscribe the user
	if err := s.subRepo.SubscribeUser(id); err != nil {
		return err
	}

	// sends a email
	// WIP

	s.log.Infoln("Subscription done successfully")
	return nil
}
