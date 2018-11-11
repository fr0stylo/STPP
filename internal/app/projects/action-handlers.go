package projects

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/pkg/entities"
	. "time-logger/internal/pkg/http-wrappers"
)

func GetAllProjectsEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	entries, err := e.DBConnection.FindAll()
	if err != nil {
		return StatusError{404, fmt.Errorf("%s", "Invalid entries")}
	}

	RespondWithJson(w, http.StatusOK, entries)

	return nil
}

func GetProjectEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r := recover(); r != nil {
			RespondWithError(w, http.StatusNotFound, "Entry not found")
		}
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	params := mux.Vars(r)

	entry, err := e.DBConnection.FindById(params["id"])

	project := entry.(Project)

	if err != nil {
		return StatusError{
			http.StatusNotFound,
			fmt.Errorf("%s", "Invalid entries"),
		}
	}

	RespondWithJson(w, http.StatusOK, project)

	return nil
}

func AddProjectEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s, %s", "Invalid request payload", err),
		}
	}
	project.ID = bson.NewObjectId()

	err := e.DBConnection.Insert(project)
	if err != nil {
		return StatusError{
			http.StatusConflict,
			fmt.Errorf("%s", err.Error()),
		}
	}
	RespondWithJson(w, http.StatusCreated, project)

	return nil
}

func UpdateProjectEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		return StatusError{http.StatusBadRequest, fmt.Errorf("%s", "Invalid request payload")}
	}

	if err := e.DBConnection.Update(project); err != nil {
		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s, %s", "Invalid request payload", err)}
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}

func DeleteProjectEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		return StatusError{http.StatusBadRequest, fmt.Errorf("%s ", "Invalid payload")}
	}

	if err := e.DBConnection.Delete(project); err != nil {
		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s ", err)}
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}
