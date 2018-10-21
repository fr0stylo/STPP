package tasks

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/", GetAllTasksEndPoint).Methods("GET")
	r.HandleFunc("/", AddTaskEndPoint).Methods("POST")
	r.HandleFunc("/", UpdateTaskEndPoint).Methods("PUT")
	r.HandleFunc("/", DeleteTaskEndPoint).Methods("DELETE")
	r.HandleFunc("/{id}", GetTaskEndPoint).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	stringedPort := ":"+ strconv.Itoa(port)

	if err := http.ListenAndServe(stringedPort, n); err != nil {
		log.Fatal(err)
	}
}
