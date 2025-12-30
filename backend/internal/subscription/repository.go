package subscription

import "github.com/rafaeldepontes/gopher-sub/internal/auth/model"

type Repository interface {
	SubscribeUser(id int64) error
	FindById(id int64) (*model.User, error)
}
