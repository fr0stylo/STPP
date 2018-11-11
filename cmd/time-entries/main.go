package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	. "time-logger/internal/app/time-entries"
	"time-logger/internal/pkg/initialization"
)

func main() {
	portPtr := flag.Int("port", 3000, "Port number")
	flag.Parse()

	env := initialization.StartupEnv()
	r := mux.NewRouter()

	s := StartServer(r, env)

	stringedPort := ":" + strconv.Itoa(*portPtr)

	if err := http.ListenAndServe(stringedPort, s); err != nil {
		log.Fatal(err)
	}
}
