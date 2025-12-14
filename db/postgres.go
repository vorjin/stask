// Package postgres is used for postgres db operations
package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	psqlinfo := os.Getenv("DB_CONN_STR")

	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Succesfully connected to postgres database")
	return db, nil
}
