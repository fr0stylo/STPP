package tasks

import (
	"github.com/codegangsta/negroni"
	"time-logger/internal/app/tasks/database"
	"time-logger/internal/pkg/server"
	. "time-logger/shared/http-wrappers"
)

func StartServer(r server.Router, env *Env) *negroni.Negroni {
	env.DBConnection = &database.TaskDAO{DB: env.DB}

	r.Handle("/", Handler{env, GetAllTasksEndPoint}).Methods("GET")
	r.Handle("/", Handler{env, AddTaskEndPoint}).Methods("POST")
	r.Handle("/", Handler{env, UpdateTaskEndPoint}).Methods("PUT")
	r.Handle("/", Handler{env, DeleteTaskEndPoint}).Methods("DELETE")
	r.Handle("/{id}", Handler{env, GetTaskEndPoint}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
