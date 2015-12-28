package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"github.com/shwoodard/jsonapi"
	"log"
	"net/http"
	"time"
)

var todos []*Todo
var serial int = 0

var listen string

type Todo struct {
	Id    int    `jsonapi:"primary,todos"`
	Title string `jsonapi:"attr,title"`
	Body  string `jsonapi:"attr,body"`
	Done  bool   `jsonapi:"attr,done"`
}

func init() {
	flag.StringVar(&listen, "port", "8080", "The port on which this application should listen for connections")
	flag.Parse()
}

func main() {
	r := httprouter.New()
	r.HandlerFunc("POST", "/todos", Logger(AddTodo, "AddTodo"))
	r.HandlerFunc("GET", "/todos", Logger(ListTodos, "ListTodos"))

	log.Printf("Awaiting connections on port %s ...", listen)

	log.Fatal("ListenAndServe: ", http.ListenAndServe(":"+listen, r))
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
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

func ListTodos(w http.ResponseWriter, r *http.Request) {
	todoInterface := make([]interface{}, len(todos))

	for i, todo := range todos {
		todoInterface[i] = todo
	}

	w.WriteHeader(http.StatusFound)
	w.Header().Set("Content-Type", "application/vnd.api+json")

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

func Logger(inner http.HandlerFunc, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
