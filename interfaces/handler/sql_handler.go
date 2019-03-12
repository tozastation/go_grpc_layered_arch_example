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
func OpenDBConnection() error {
	vender := os.Getenv("VENDER")
	connectionString := os.Getenv("CONNECTION_STRING")
	DB, err := sql.Open(vender, connectionString)
	err = DB.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Establised DB Connection!!")
	return nil
}
