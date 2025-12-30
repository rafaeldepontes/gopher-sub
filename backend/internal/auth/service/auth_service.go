package service

import (
	"time"

	"github.com/rafaeldepontes/gopher-sub/internal/auth"
	"github.com/rafaeldepontes/gopher-sub/internal/auth/model"
	"github.com/rafaeldepontes/gopher-sub/internal/auth/repository"
	msg "github.com/rafaeldepontes/gopher-sub/internal/errors-messages"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const Cost = 16

type authService struct {
	authRepo auth.Repository
	log      *logrus.Logger
}

func NewService() auth.Service {
	return &authService{
		authRepo: repository.NewRepository(),
		log:      logger.GetLogger(),
	}
}

// Login implements [auth.Service].
func (s *authService) Login(userReq *model.UserReq) error {
	if userReq.Email == "" {
		return msg.ErrInvalidEmail
	}

	if userReq.Password == "" || len(userReq.Password) <= 4 {
		return msg.ErrInvalidPassword
	}

	s.log.Infof("Log in attempt to email: %s\n", userReq.Email)

	user, err := s.authRepo.FindByEmail(userReq.Email)
	if err != nil {
		return msg.ErrUserNotFound
	}

	if err = bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(userReq.Password),
	); err != nil {
		return msg.ErrInvalidCredentials
	}
	return nil
}

// Register implements [auth.Service].
func (s *authService) Register(userReq *model.UserReq) error {
	if userReq.Email == "" {
		return msg.ErrInvalidEmail
	}

	if userReq.Password == "" || len(userReq.Password) <= 4 {
		return msg.ErrInvalidPassword
	}

	s.log.Infof("Registering a new user, email: %s\n", userReq.Email)

	user_, err := s.authRepo.FindByEmail(userReq.Email)
	if user_ != nil && err == nil {
		return msg.ErrUserAlreadyExist
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), Cost)
	if err != nil {
		return msg.ErrInternalServer
	}

	user := model.User{
		Email:          userReq.Email,
		HashedPassword: string(hp),
		CreatedAt:      time.Now(),
		IsSubscribed:   false,
	}
	if err := s.authRepo.Register(&user); err != nil {
		return err
	}

	return nil
}
