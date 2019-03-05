package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
)

// OpenDBConnection is ...
func OpenDBConnection() *sql.DB {
	vender := os.Getenv("VENDER")
	connectionString := os.Getenv("CONNECTION_STRING")
	db, err := sql.Open(vender, connectionString)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Establised DB Connection!!")
	return db
}
