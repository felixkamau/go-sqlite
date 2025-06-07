package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//  Create a new db  and create a table in the db
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// defer close() ensures the resource is closed when main() finishes
	defer db.Close()
	// sql-statement to create a table
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			task TEXT,
			status BOOL
		);
	`
	// Execute the sql-script
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	// `sqlite3 tasks.db .schema` to check schema
	log.Println("Table 'tasks' created successfully")
}
