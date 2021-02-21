package server

import (
	"net/http"

	"github.com/go-run/todolist_ex/db"
)

type Config struct {
	Address     string
	DatabaseURL string
}

func ListenAndServe(c Config) error {
	if err := db.Connect(c.DatabaseURL); err != nil {
		return err
	}
	http.ListenAndServe(c.Address,
		loggingMiddleware(api.TodolistAPI()))
}
