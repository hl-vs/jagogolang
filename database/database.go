package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {

	// open database
	fmt.Println("[DATABASE] Opening ...")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// test connection
	fmt.Println("[DATABASE] PING database ...")
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// set connection pool settings (optional but recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	fmt.Println("[DATABASE] Connected succesfully.")
	return db, nil
}
