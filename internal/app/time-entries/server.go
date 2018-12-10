package time_entries

import (
	"github.com/codegangsta/negroni"
	"time-logger/internal/app/tasks/database"
	"time-logger/internal/pkg/server"
	"time-logger/shared/http-wrappers"
)

func StartServer(r server.Router, env *http_wrappers.Env) *negroni.Negroni {
	env.DBConnection = &database.TaskDAO{DB: env.DB}

	//r.Handle("/", http_wrappers.Handler{env, GetAllTimeEntriesEndPoint}).Methods("GET")
	r.Handle("/", http_wrappers.Handler{env, AddTimeEntryEndPoint}).Methods("POST")
	r.Handle("/", http_wrappers.Handler{env, UpdateTimeEntryEndPoint}).Methods("PUT")
	r.Handle("/", http_wrappers.Handler{env, DeleteTimeEntryEndPoint}).Methods("DELETE")
	//r.Handle("/{id}", http_wrappers.Handler{env, GetTimeEntryByIdEndPoint}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
