package email

import "github.com/rafaeldepontes/gopher-sub/internal/auth/model"

type Service interface {
	SendMail(user *model.User) error
}
