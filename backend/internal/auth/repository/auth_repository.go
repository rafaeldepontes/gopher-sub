package repository

import (
	"database/sql"

	"github.com/rafaeldepontes/gopher-sub/internal/auth"
	"github.com/rafaeldepontes/gopher-sub/internal/auth/model"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
	"github.com/rafaeldepontes/gopher-sub/pkg/database/postgres"
	"github.com/sirupsen/logrus"
)

type authRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewRepository() auth.Repository {
	return &authRepository{
		db:  postgres.GetDb(),
		log: logger.GetLogger(),
	}
}

// FindByEmail implements [auth.Repository].
func (r *authRepository) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, hashed_password, created_at FROM users u where u.email = $1
	`
	var user model.User
	if err := r.db.QueryRow(
		query,
		email,
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

// Register implements [auth.Repository].
func (r *authRepository) Register(user *model.User) error {
	query := `
		INSERT INTO users (email, hashed_password, is_subscribed, created_at) VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(query, user.Email, user.HashedPassword, user.IsSubscribed, user.CreatedAt)
	if err != nil {
		r.log.Errorf("couldn't insert user: %s because %v\n", user.Email, err)
		return err
	}
	return nil
}
