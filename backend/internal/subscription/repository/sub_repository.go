package repository

import (
	"database/sql"

	"github.com/rafaeldepontes/gopher-sub/internal/auth/model"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/internal/subscription"
	"github.com/rafaeldepontes/gopher-sub/pkg/database/postgres"
	"github.com/sirupsen/logrus"
)

type subRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewRepository() subscription.Repository {
	return &subRepository{
		db:  postgres.GetDb(),
		log: logger.GetLogger(),
	}
}

// SubscribeUser implements [subscription.Repository].
func (r *subRepository) SubscribeUser(id int64) error {
	query := `
		UPDATE users
		SET is_subscribed = true
		WHERE id = $1
	`
	result, err := r.db.Exec(query, id)
	if err != nil {
		r.log.Errorf("couldn't update user: %d because %v\n", id, err)
		return err
	}
	qtdAffected, _ := result.RowsAffected()

	r.log.Infoln("rows affected, ", qtdAffected)
	return nil
}

// FindByEmail implements [auth.Repository].
func (r *subRepository) FindById(id int64) (*model.User, error) {
	query := `
		SELECT id, email, hashed_password, created_at FROM users u where u.id = $1
	`
	var user model.User
	if err := r.db.QueryRow(
		query,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
	); err != nil {
		r.log.Errorln("couldn't find user, ", err)
		return nil, err
	}

	return &user, nil
}
