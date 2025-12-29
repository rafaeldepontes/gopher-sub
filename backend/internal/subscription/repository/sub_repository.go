package repository

import (
	"database/sql"

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
