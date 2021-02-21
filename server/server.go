package server

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist/api"
	"github.com/jeonjonghyeok/todolist/db"
)

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(c Config) error {
	if err := db.Connect(c.DatabaseURL); err != nil {
		return err
	}

	return http.ListenAndServe(c.Address,
		loggingMiddleware(api.TodoListAPI()))
}
