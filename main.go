package main

import (
	"log"

	"github.com/jeonjonghyeok/todolist/server"
)

func main() {
	if err := server.ListenAndServe(server.Config{
		Address:     ":8082",
		DatabaseURL: "postgres://postgres:tododbpasswd123@tododb.learningspoons.danslee.com:5432/postgres?sslmode=require",
	}); err != nil {
		log.Fatalln(err)
	}
}
