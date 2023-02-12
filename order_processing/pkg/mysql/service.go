package mysql

import (
	"database/sql"
	"log"
)

type Service struct {
	// database connection
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	if db == nil {
		log.Fatalln("NewService: Null pointer error")
		return nil
	}

	service := Service{
		db: db,
	}

	return &service
}
