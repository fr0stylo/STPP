package time_entries

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/app/time-entries/database"
	. "time-logger/internal/pkg/config"
	. "time-logger/internal/pkg/dtos"
	. "time-logger/internal/pkg/entities"
	. "time-logger/internal/pkg/http-wrappers"
)

var dao TimeEntryDAO

func init() {
	r := Reader{}
	config := r.GetConfig()
	dao.Server = config.Server
	dao.Database = config.Database

	dao.Connect()
}

func GetAllTimeEntriesEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entries []TimeEntry
	var err error

	entries, err = dao.FindAll()

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	var dtos []TimeEntryDTO
	chProj  := make(chan HTTPResponse)
	chTask  := make(chan HTTPResponse)

	go MakeRequest("http://project-service:3000/", chProj)
	go MakeRequest("http://tasks-service:3000/", chTask)

	var projects []Project
	var tasks []Project

	if json.NewDecoder(bytes.NewReader((<-chProj).Body)).Decode(&projects); err != nil {
		RespondWithError(w, http.StatusFailedDependency, "Network Error Occoured, "+ err.Error())
	}
	if json.NewDecoder(bytes.NewReader((<-chTask).Body)).Decode(&tasks); err != nil {
		RespondWithError(w, http.StatusFailedDependency, "Network Error Occoured, "+ err.Error())
	}

	projectMap := make(map[bson.ObjectId]string)
	for _, p := range projects {
		projectMap[p.ID] = p.Name
	}

	tasksMap := make(map[bson.ObjectId]string)
	for _, t := range tasks {
		tasksMap[t.ID] = t.Name
	}


	for _, u := range entries {
		dto := TimeEntryDTO{
			Description: u.Description,
			ID: u.ID,
			Project: projectMap[u.Project],
			ProjectId: u.Project,
			TaskId: u.Task,
			Task: tasksMap[u.Task],
			TimeSpent: u.TimeSpent,
		}

		dtos = append(dtos, dto)
	}

	RespondWithJson(w, http.StatusOK, dtos)
}

func GetTimeEntryByIdEndPoint(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			RespondWithError(w, http.StatusNotFound, "Not Found")
		}
		r.Body.Close()
	}()
	params := mux.Vars(r)

	entry, err := dao.FindById(params["id"])

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Invalid entries")
		return
	}

	response, err := http.Get(string("http://project-service:3000/" + entry.Project.Hex()))

	if err != nil {
		RespondWithError(w, http.StatusFailedDependency, "Network Error Occoured, " + err.Error())
		return
	}

	project := Project{}
	defer response.Body.Close()
	if json.NewDecoder(response.Body).Decode(&project); err != nil {
		RespondWithError(w, http.StatusFailedDependency, "Network Error Occoured, "+ err.Error())
	}

	dto := TimeEntryDTO{
		Description: entry.Description,
		ID: entry.ID,
		Project: project.Name,
		ProjectId: entry.Project,
		TaskId: entry.Task,
		TimeSpent: entry.TimeSpent,
	}

	RespondWithJson(w, http.StatusOK, dto)
}

func AddTimeEntryEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var timeEntry TimeEntry
	if err := json.NewDecoder(r.Body).Decode(&timeEntry); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload, " + err.Error())
		return
	}
	timeEntry.ID = bson.NewObjectId()

	if err := dao.Insert(timeEntry); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusCreated, timeEntry)
}

func UpdateTimeEntryEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entry TimeEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(entry); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteTimeEntryEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var entry TimeEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(entry); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

