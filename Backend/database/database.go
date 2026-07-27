package database


import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

const schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

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

	// 1. Create the tables in memory
	db.MustExec(schema)

	// 2. Optional: Seed mock data so your UI isn't completely empty when testing
	seedData := `INSERT INTO users (name) VALUES ('Darien'), ('Miller'), ('Vicky');`
	db.MustExec(seedData)

	DB = db
}