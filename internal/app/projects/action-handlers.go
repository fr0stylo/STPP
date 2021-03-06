package projects

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/pkg/entities"
	. "time-logger/shared/http-wrappers"
)

func GetAllProjectsEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	results, err := e.DBConnection.FindAll()

	if err != nil {
		return StatusError{404, fmt.Errorf("%s", "Invalid entries")}
	}

	RespondWithJson(w, http.StatusOK, results)

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

	if err != nil {
		return StatusError{
			http.StatusNotFound,
			fmt.Errorf("%s", "Invalid entries"),
		}
	}

	RespondWithJson(w, http.StatusOK, entry)

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
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		return StatusError{http.StatusBadRequest, fmt.Errorf("%s", "Invalid request payload")}
	}
	defer r.Body.Close()

	if err := e.DBConnection.Update(project); err != nil {
		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s, %s", "Invalid request payload", err)}
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}

func DeleteProjectEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	if err := e.DBConnection.Delete(params["id"]); err != nil {
		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s ", err)}
	}

	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}
