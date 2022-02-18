package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(dbConfig DBConfig) (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
	)

	database, err := sql.Open("postgres", connString)

	if err != nil {
		return database, err
	}

	return database, nil
}
