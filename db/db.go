package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	Pet  *PetStorage
	User *UserStorage
}

func NewDB(connStr string) (*sql.DB, error) {

	// connStr := "user=postgres dbname=postgres password=petDBsecr3t sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// pg := &DB{
	// 	db: db,
	// }

	return db, nil
}
