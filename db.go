package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func initDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func addTask(db *sql.DB, description string) {
	_, err := db.Exec("INSERT INTO tasks (description) VALUES (?)", description)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✔️  Task added successfully!")
}

func listTasks(db *sql.DB) {
	rows, err := db.Query("SELECT id, description, completed FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n--- Task List ---")
	fmt.Println("ID | Description          | Status")
	fmt.Println("---|----------------------|--------")

	for rows.Next() {
		var id int
		var description string
		var completed bool
		rows.Scan(&id, &description, &completed)
		status := "❌ Incomplete"
		if completed {
			status = "✅ Completed"
		}
		fmt.Printf("%-3d| %-21s | %s\n", id, description, status)
	}
}

func completeTask(db *sql.DB, id int) {
	result, err := db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("❌ Task not found.")
	} else {
		fmt.Println("✔️  Task marked as completed!")
	}
}
