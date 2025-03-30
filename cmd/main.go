package main

import (
	"log"
	"net/http"

	"github.com/alberto-debug/todo-backend/internal/handlers"
	"github.com/alberto-debug/todo-backend/pkg/db"
)

func main() {
	db.InitDB()
	http.HandleFunc("/api/todos", handlers.TodoHandler)
	log.Println("Server started at :8080, test qand enjoy your app")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
