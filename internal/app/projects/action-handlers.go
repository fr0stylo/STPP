package projects

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"

	. "time-logger/internal/app/projects/database"
	. "time-logger/internal/pkg/config"
	. "time-logger/internal/pkg/entities"
	. "time-logger/internal/pkg/http-wrappers"
)

var dao ProjectDAO

func init() {
	config := GetConfig()
	dao.Server = config.Server
	dao.Database = config.Database

	dao.Connect()
}

func GetAllProjectsEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	entries, err := dao.FindAll()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	RespondWithJson(w, http.StatusOK, entries)
}

func GetProjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer func () {
		if r := recover(); r != nil {
			RespondWithError(w, http.StatusNotFound, "Entry not found")
		}
		r.Body.Close()
	}()
	params := mux.Vars(r)

	entry, err := dao.FindById(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	RespondWithJson(w, http.StatusOK, entry)
}

func AddProjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		log.Fatal(err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	project.ID = bson.NewObjectId()

	if err := dao.Insert(project); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, project)
}

func UpdateProjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(project); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteProjectEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(project); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
