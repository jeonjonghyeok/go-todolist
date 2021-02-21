package api

import (
	"net/http"

	"github.com/jeonjonghyeok/go-run/todolist_ex/db"
)

func getTodoLists(w http.ResponseWriter, r *http.Request) {
	lists, err := db.GetTodoList()
	must(err)
	writeJSON(w, lists)

}
