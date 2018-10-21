package tasks

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/app/tasks/database"
	. "time-logger/internal/pkg/config"
	. "time-logger/internal/pkg/entities"
	. "time-logger/internal/pkg/http-wrappers"
)

var dao TaskDAO

func init() {
	config := GetConfig()
	dao.Server = config.Server
	dao.Database = config.Database

	dao.Connect()
}

func GetAllTasksEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	entries, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	RespondWithJson(w, http.StatusOK, entries)
}

func GetTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)

	entry, err := dao.FindById(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	RespondWithJson(w, http.StatusOK, entry)
}

func AddTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	task.ID = bson.NewObjectId()

	if err := dao.Insert(task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, task)
}

func UpdateTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteTaskEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(task); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

