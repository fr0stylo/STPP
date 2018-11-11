package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/pkg/entities"
	. "time-logger/internal/pkg/http-wrappers"
)

func GetAllTasksEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	entries, err := e.DBConnection.FindAll()
	if err != nil {
		return StatusError{
			http.StatusInternalServerError,
			fmt.Errorf("%s", "Invalid entries"),
		}
	}

	RespondWithJson(w, http.StatusOK, entries)

	return nil
}

func GetTaskEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	params := mux.Vars(r)

	entry, err := e.DBConnection.FindById(params["id"])

	if err != nil {
		return StatusError{
			http.StatusInternalServerError,
			fmt.Errorf("%s", "Invalid entries"),
		}
	}

	RespondWithJson(w, http.StatusOK, entry)

	return nil
}

func AddTaskEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", "Invalid request payload"),
		}
	}
	task.ID = bson.NewObjectId()

	if err := e.DBConnection.Insert(task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", err.Error()),
		}
	}
	RespondWithJson(w, http.StatusCreated, task)

	return nil
}

func UpdateTaskEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", "Invalid request payload"),
		}
	}

	if err := e.DBConnection.Update(task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", err.Error()),
		}
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}

func DeleteTaskEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", "Invalid request payload"),
		}
	}

	if err := e.DBConnection.Delete(task); err != nil {
		return StatusError{
			http.StatusBadRequest,
			fmt.Errorf("%s", err.Error()),
		}
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}
