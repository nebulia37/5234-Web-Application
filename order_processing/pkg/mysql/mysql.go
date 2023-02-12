package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect connects to a MySQL database
func Connect(user string, pass string, dbname string, address string) (*sql.DB, error) {
	// build the dsn string
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user, pass, address, dbname)

	// connect to database
	db, err := sql.Open("mysql", dsn)
	if err != nil || db == nil {
		return nil, fmt.Errorf("connect dsn[%s]: %v", dsn, err)
	}

	// test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("connect dsn[%s]: %v", dsn, err)
	}

	log.Printf("Connected to database at [%s]/[%s]", address, dbname)

	return db, nil
}
