package auth

import "github.com/rafaeldepontes/gopher-sub/internal/auth/model"

type Service interface {
	Login(user *model.UserReq) error
	Register(user *model.UserReq) error
}
