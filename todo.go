package main

import (
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/julienschmidt/httprouter"
	"github.com/shwoodard/jsonapi"
	"log"
	"net/http"
	"strings"
	"time"
)

var db bolt.DB
var todos []*Todo
var serial int = 0

var listen string

type Todo struct {
	Id    int    `json:"id" jsonapi:"primary,todos"`
	Title string `json:"title" jsonapi:"attr,title"`
	Body  string `json:"body" jsonapi:"attr,body"`
	Done  bool   `json:"done" jsonapi:"attr,done"`
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func init() {
	flag.StringVar(&listen, "port", "8080", "The port on which this application should listen for connections")
	flag.Parse()
}

func main() {
	transaction(func(db *bolt.DB) error {
		return db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte("todos"))
			if err != nil {
				return fmt.Errorf("Bucket creation: %s", err)
			}
			return nil
		})
	})

	r := httprouter.New()
	addRoutes(r)

	log.Printf("Awaiting connections on port %s ...", listen)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":"+listen, r))
}

func addRoutes(r *httprouter.Router) {
	routes := map[string]Route{
		"AddTodo": {
			Method:  "POST",
			Path:    "/todos",
			Handler: AddTodo,
		},
		"ListTodos": {
			Method:  "GET",
			Path:    "/todos",
			Handler: ListTodos,
		},
	}

	options := make(map[string][]string)
	for name, route := range routes {
		r.HandlerFunc(route.Method, route.Path, Logger(Headers(route.Handler), name))
		options[route.Path] = append(options[route.Path], route.Method)
	}

	for path, methods := range options {
		methods = append(methods, "OPTIONS")
		r.HandlerFunc("OPTIONS", path, http.HandlerFunc(Logger(Headers(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ", "))
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
		}), "Options("+path+")")))
	}
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

	if err := jsonapi.MarshalOnePayload(w, todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := getTodos()
	if err != nil {
		log.Printf("Getting todos: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalManyPayload(w, todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getTodos() ([]interface{}, error) {
	todos := make([]interface{}, 0)

	err := transaction(func(db *bolt.DB) error {
		return db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("todos"))
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				todo := &Todo{}
				if err := json.Unmarshal(v, todo); err != nil {
					return err
				}
				todos = append(todos, todo)
			}

			return nil
		})
	})

	return todos, err
}

func saveTodo(todo *Todo) error {
	return transaction(func(db *bolt.DB) error {
		return db.Update(func(tx *bolt.Tx) error {
			// Retrieve the users bucket.
			// This should be created when the DB is first opened.
			b := tx.Bucket([]byte("todos"))

			// Generate ID for the user.
			// This returns an error only if the Tx is closed or not writeable.
			// That can't happen in an Update() call so I ignore the error check.
			id, _ := b.NextSequence()
			todo.Id = int(id)

			// Marshal user data into bytes.
			buf, err := json.Marshal(todo)
			if err != nil {
				return err
			}

			// Persist bytes to users bucket.
			return b.Put(itob(todo.Id), buf)
		})
	})
}

func Headers(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/vnd.api+json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		inner(w, r)
	})
}

func Logger(inner http.HandlerFunc, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner(w, r)

		log.Printf(
			"%-15s%-15s%-30s%s",
			r.RequestURI,
			r.Method,
			name,
			time.Since(start),
		)
	})
}

func transaction(tx func(*bolt.DB) error) error {
	db, err := bolt.Open("todo.db", 0600, &bolt.Options{Timeout: 3 * time.Second})
	if err != nil {
		log.Fatal("BoltDB: ", err.Error())
	}
	defer db.Close()

	return tx(db)
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
