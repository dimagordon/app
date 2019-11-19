package repository

import "database/sql"

type Todo struct {
	ID    string
	Title string `json:"title"`
	Text  string `json:"title"`
}

type TodoRepo struct {
	db *sql.DB
}
