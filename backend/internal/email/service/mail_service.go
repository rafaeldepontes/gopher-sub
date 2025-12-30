package service

import (
	"net/smtp"

	"github.com/rafaeldepontes/gopher-sub/internal/auth/model"
	"github.com/rafaeldepontes/gopher-sub/internal/email"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/sirupsen/logrus"
)

type emailService struct {
	log *logrus.Logger
}

func NewService() email.Service {
	return &emailService{
		log: logger.GetLogger(),
	}
}

func (s *emailService) SendMail(user *model.User) error {
	s.log.Infoln("Sending email to", user.Email)
	from := "test@gmail.com"
	to := []string{user.Email}

	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + user.Email + "\r\n" +
			"Subject: Subscription!\r\n" +
			"\r\n" +
			"Your payment was accept! You can now enjoy your subscription!\r\n",
	)

	if err := smtp.SendMail("localhost:1025", nil, from, to, msg); err != nil {
		s.log.Errorln("sendmail:", err)
		return err
	}
	s.log.Infoln("email sent (captured by MailHog) to", user.Email)
	return nil
}
