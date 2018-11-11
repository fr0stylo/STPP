package projects

import (
	"github.com/codegangsta/negroni"
	"time-logger/internal/app/projects/database"
	. "time-logger/internal/pkg/http-wrappers"
	"time-logger/internal/pkg/server"
)

func StartServer(r server.Router, env *Env) *negroni.Negroni{
	env.DBConnection = &database.ProjectDAO{ DB: env.DB}

	r.Handle("/", Handler{env, GetAllProjectsEndPoint}).Methods("GET")
	r.Handle("/", Handler{env, AddProjectEndPoint}).Methods("POST")
	r.Handle("/", Handler{env, UpdateProjectEndPoint}).Methods("PUT")
	r.Handle("/", Handler{env, DeleteProjectEndPoint}).Methods("DELETE")
	r.Handle("/{id}", Handler{env, GetProjectEndPoint}).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}
