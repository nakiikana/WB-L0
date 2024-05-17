package repository

import (
	"fmt"
	"tools/internals/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func NewPostgresDB(config *config.Configuration) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password, config.DB.Sslmode))
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to db")
	}
	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "could not ping created db")
	}
	return db, nil
}
