package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabriel_assis7/simple-go-mod/configs"
	"github.com/gabriel_assis7/simple-go-mod/models"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db := configs.ConnectDB()
	defer db.Close()

	err = db.QueryRow("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, task.Status).Scan(&task.ID)

	if err != nil {
		http.Error(w, "Could not create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	db := configs.ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		http.Error(w, "Could not fetch tasks", http.StatusInternalServerError)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			http.Error(w, "Error reading task", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := configs.ConnectDB()
	defer db.Close()

	var task models.Task
	err := db.QueryRow("SELECT id, title, description, status FROM tasks where id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, "Could not fetch task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	db := configs.ConnectDB()
	defer db.Close()

	result, err := db.Exec(
		"UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4",
		task.Title, task.Description, task.Status, id,
	)

	if err != nil {
		http.Error(w, "Could not update task", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task.ID = intID

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := configs.ConnectDB()
	defer db.Close()

	result, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
