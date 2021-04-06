package pkg

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // ...
)

type Connection struct {
	databaseURI string
	DB          *sqlx.DB
}

func NewConnection(uri string) *Connection {
	return &Connection{
		databaseURI: uri,
	}
}

func (conn *Connection) Open() error {
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "postgres", conn.databaseURI)
	if err != nil {
		return err
	}

	conn.DB = db

	return nil
}