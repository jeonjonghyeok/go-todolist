package api

import (
	"net/http"

	"github.com/jeonjonghyeok/todolist/db"
	"github.com/jeonjonghyeok/todolist/todo"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoLists()
	must(err)
	writeJSON(w, lists)

}
func createTodoList(w http.ResponseWriter, r *http.Request) {
	var req todo.List
	parseJSON(r.Body, &req)
	todoList, err := db.CreateTodoList(req.Name)
	must(err)
	writeJSON(w, todoList)
}
