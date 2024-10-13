package main

import (
    "fmt"
    "strconv"
    "strings"
)

func showMenu() {
    fmt.Println("\n--- Task Tracker ---")
    fmt.Println("Commands:")
    fmt.Println("- add: Add a new task")
    fmt.Println("- list: List all tasks")
    fmt.Println("- complete: Mark a task as completed")
    fmt.Println("- exit: Exit the program")
}

func main() {
    db := initDB()
    defer db.Close()

    showMenu()

    for {
        fmt.Print("\nEnter command: ")
        var input string
        fmt.Scanln(&input)
        input = strings.ToLower(input)

        switch input {
        case "add":
            fmt.Print("Enter task description: ")
            var description string
            fmt.Scanln(&description)
            addTask(db, description)

        case "list":
            listTasks(db)

        case "complete":
            fmt.Print("Enter task ID to complete: ")
            var idStr string
            fmt.Scanln(&idStr)
            id, err := strconv.Atoi(idStr)
            if err != nil {
                fmt.Println("❌ Invalid task ID.")
                continue
            }
            completeTask(db, id)

        case "exit":
            fmt.Println("Goodbye!")
            return

        case "help":
            showMenu()

        default:
            fmt.Println("❌ Invalid command. Type 'help' for the list of commands.")
        }
    }
}
