package repository

import (
	"context"
	"database/sql"
)

type Todo struct {
	ID    string
	Title string `json:"title"`
	Text  string `json:"title"`
}

type TodoRepo struct {
	db *sql.DB
}

func New(db *sql.DB) *TodoRepo {
	return &TodoRepo{db}
}

func (t *TodoRepo) Create(ctx context.Context) {
	//u := model.Todo{
	//	Title: title,
	//	Text: text
	//}

}
