package config

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/", GetConfig).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	stringedPort := ":"+ strconv.Itoa(port)

	if err := http.ListenAndServe(stringedPort, n); err != nil {
		log.Fatal("error: ", err)
	}
}
