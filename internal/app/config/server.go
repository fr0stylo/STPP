package config

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	. "time-logger/internal/pkg/dtos"
)

type Server struct {
	r mux.Router

}

func StartServer(port int) {
	r := mux.NewRouter()
	c := Config{}
	cr := ConfigReader{
		c,
	}

	r.HandleFunc("/", cr.GetConfig).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	stringedPort := ":"+ strconv.Itoa(port)

	if err := http.ListenAndServe(stringedPort, n); err != nil {
		log.Fatal("error: ", err)
	}
}
