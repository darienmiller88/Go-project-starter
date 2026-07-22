package database


import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

// Extend this to use an actual SQL file, or Postgres
func ConnectToSQL() {
	db, err := sqlx.Connect("sqlite", ":memory:")

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	
	if err != nil {
		fmt.Println("db connection fail:", err)
	}else{
		fmt.Println("Connection established! :)")
	}

	DB = db
}