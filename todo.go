package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/shwoodard/jsonapi"
	"log"
	"net/http"
)

var todos []*Todo
var serial int = 0

type Todo struct {
	Id    int    `jsonapi:"primary,todos"`
	Title string `jsonapi:"attr,title"`
	Body  string `jsonapi:"attr,body"`
	Done  bool   `jsonapi:"attr,done"`
}

func main() {
	r := httprouter.New()
	r.POST("/todos", AddTodo)
	r.GET("/todos", ListTodos)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func AddTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := new(Todo)

	if err := jsonapi.UnmarshalPayload(r.Body, todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := saveTodo(todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/vnd.api+json")

	if err := jsonapi.MarshalOnePayload(w, todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todoInterface := make([]interface{}, len(todos))

	for i, todo := range todos {
		todoInterface[i] = todo
	}

	if err := jsonapi.MarshalManyPayload(w, todoInterface); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func saveTodo(todo *Todo) error {
	todo.Id = serial
	serial++
	todos = append(todos, todo)
	return nil
}
