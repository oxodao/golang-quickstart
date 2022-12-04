package orm

import (
	"github.com/jmoiron/sqlx"
	"github.com/oxodao/go-quickstart/config"

	_ "github.com/mattn/go-sqlite3"
)

var GET *ORM

type ORM struct {
	db *sqlx.DB
}

func Load() error {
	db := sqlx.MustConnect("sqlite3", config.GET.Database.Path)

	GET = &ORM{
		db: db,
	}

	return nil
}
