package projects

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/", GetAllProjectsEndPoint).Methods("GET")
	r.HandleFunc("/", AddProjectEndPoint).Methods("POST")
	r.HandleFunc("/", UpdateProjectEndPoint).Methods("PUT")
	r.HandleFunc("/", DeleteProjectEndPoint).Methods("DELETE")
	r.HandleFunc("/{id}", GetProjectEndPoint).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	stringedPort := ":"+ strconv.Itoa(port)

	if err := http.ListenAndServe(stringedPort, n); err != nil {
		log.Fatal(err)
	}
}
