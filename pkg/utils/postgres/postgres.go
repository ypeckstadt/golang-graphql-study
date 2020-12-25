package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// New creates a new Postgres database connection
func New(connection ConnectionSettings) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connection.Host, connection.Port, connection.Username, connection.Password, connection.Name,
	)
	return sql.Open("postgres", psqlInfo)
}
