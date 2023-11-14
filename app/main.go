package main

import (
	"log"
	"net/http"
	"todo-app/internal/app"
	internalhttp "todo-app/pkg/http"
)

func main() {
	todoMap := make(app.ToDoMap)
	internalhttp.HandleToDoRequests(todoMap)

	log.Println("Server is running on http://localhost:8091")
	log.Fatal(http.ListenAndServe(":8091", nil))
}
