package auth

import "github.com/rafaeldepontes/gopher-sub/internal/auth/model"

type Repository interface {
	FindByEmail(email string) (*model.User, error)
	Register(user *model.User) error
}
