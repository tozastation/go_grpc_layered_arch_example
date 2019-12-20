package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
)

// DB is .
var DB *sql.DB

// OpenDBConnection is ...
func OpenDBConnection() (*sql.DB, error) {
	vender := os.Getenv("VENDER")
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open(vender, connectionString)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Establised DB Connection!!")
	return db, nil
}
