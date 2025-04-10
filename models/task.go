package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

var CreateTableSQL = `
CREATE TABLE IF NOT EXISTS tasks (
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	status BOOLEAN NOT NULL DEFAULT false
);`
