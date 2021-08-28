package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	CertPath string
}

func NewDB(addr string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", addr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func GenerateAddr(connection *Connection) string {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s",
		connection.Host,
		connection.User,
		connection.Password,
		connection.DBName,
		connection.Port,
		connection.SSLMode,
		connection.CertPath,
	)
	return connStr
}
