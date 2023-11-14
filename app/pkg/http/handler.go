package http

import (
	"encoding/json"
	"io"
	"net/http"
	"todo-app/internal/app"
)

func HandleToDoRequests(todoMap app.ToDoMap) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			title := r.FormValue("title")
			todoMap.AddItem(title)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, `{"alive": true}`)
		}
	})

	http.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			items := todoMap.GetAllItems()
			json.NewEncoder(w).Encode(items)
		}
	})
}
