package handler

import (
	"net/http"

	"github.com/alberto-debug/todo-backend/models"
	"github.com/alberto-debug/todo-backend/pkg/db"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	todos := []models.Todo{}
}
