package postgres

import (
	"database/sql"
	"fmt"
	"github.com/alihaqberdi/goga_go/internal/config"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func InitDB() error {
	newDB, err := sql.Open("postgres", config.POSTGRES_URI)
	if err != nil {
		return fmt.Errorf("failed to connect to the sql_db: %s", err)
	}

	// Test the connection
	err = newDB.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping sql_db:%s ", err)
	}

	db = newDB
	return nil
}

func CloseDB() error {
	return db.Close()
}
