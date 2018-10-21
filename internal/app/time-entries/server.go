package time_entries

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func StartServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/", GetAllTimeEntriesEndPoint).Methods("GET")
	r.HandleFunc("/", AddTimeEntryEndPoint).Methods("POST")
	r.HandleFunc("/", UpdateTimeEntryEndPoint).Methods("PUT")
	r.HandleFunc("/", DeleteTimeEntryEndPoint).Methods("DELETE")
	r.HandleFunc("/{id}", GetTimeEntryByIdEndPoint).Methods("GET")


	n := negroni.Classic()
	n.UseHandler(r)

	stringedPort := ":"+ strconv.Itoa(port)

	if err := http.ListenAndServe(stringedPort, n); err != nil {
		log.Fatal(err)
	}
}