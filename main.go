package main

import (
	"log"

	"github.com/jeonjonghyeok/go-run/todolist_ex/server"
)

func main() {
	if err := server.ListenAndServe(
		serv.Config{
			Address:     ":8080",
			DatabaseURL: "postgres://postgres:tododbpasswd123@tododb.learningspoons.danslee.com:5432/postgres?sslmode=require",
		}); err != nil {
		log.Fatalln(err)
	}
}
