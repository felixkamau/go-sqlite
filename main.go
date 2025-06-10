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
	id := 1
	result, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %d rows(s) \n", affectedRows)
	// rows, err := db.Query("SELECT id, task, status FROM tasks")
	// sql-statement to create a table
	// sqlStmt := `
	// 	CREATE TABLE IF NOT EXISTS tasks (
	// 		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 		task TEXT,
	// 		status BOOL
	// 	);
	// `
	// // Execute the sql-script
	// _, err = db.Exec(sqlStmt)

	// Insert a new task
	// task := "Learn go"
	// status := true // or false if not done
	// _, err = db.Exec("INSERT INTO tasks(task, status) VALUES(?, ?)", task, status)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int
	// 	var task string
	// 	var status bool
	// 	err = rows.Scan(&id, &task, &status)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("Task: %d, %s, %v", id, task, status)

	// }

	// if err = rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// `sqlite3 tasks.db .schema` to check schema
	// log.Println("New 'task' added successfully")

	// with db created we can perform CRUD ops
}
