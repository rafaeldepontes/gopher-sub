package postgres

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rafaeldepontes/gopher-sub/internal/logger"
)

var db *sql.DB

func GetDb() *sql.DB {
	if db != nil {
		return db
	}
	_ = initConn()
	return db
}

func Close() error {
	return db.Close()
}

func initConn() error {
	database, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.GetLogger().Errorln("something went wrong connection to the database\n[Error]", err)
		return err
	}

	if err = database.Ping(); err != nil {
		logger.GetLogger().Errorln("something went wrong while checking the database connection\n[Error]", err)
		return err
	}

	db = database
	return nil
}
