package time_entries

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	. "time-logger/internal/pkg/entities"
	. "time-logger/shared/http-wrappers"
)

//func GetAllTimeEntriesEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
//	defer func() {
//		if r.Body != nil {
//			r.Body.Close()
//		}
//	}()
//
//	var entries []TimeEntry
//
//	interfaces, err := e.DBConnection.FindAll()
//
//	entries, _ = interfaces.([]TimeEntry)
//
//	if err != nil {
//		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s", "Invalid entries")}
//	}
//
//	var dtos []TimeEntryDTO
//	chProj := make(chan HTTPResponse)
//	chTask := make(chan HTTPResponse)
//
//	go MakeRequest("http://project-service:3000/", chProj)
//	go MakeRequest("http://tasks-service:3000/", chTask)
//
//	var projects []Project
//	var tasks []Project
//
//	if json.NewDecoder(bytes.NewReader((<-chProj).Body)).Decode(&projects); err != nil {
//		return StatusError{http.StatusFailedDependency, fmt.Errorf("%s", "Network Error Occoured")}
//	}
//	if json.NewDecoder(bytes.NewReader((<-chTask).Body)).Decode(&tasks); err != nil {
//		return StatusError{http.StatusFailedDependency, fmt.Errorf("%s", "Network Error Occoured")}
//	}
//
//	projectMap := make(map[bson.ObjectId]string)
//	for _, p := range projects {
//		projectMap[p.ID] = p.Name
//	}
//
//	tasksMap := make(map[bson.ObjectId]string)
//	for _, t := range tasks {
//		tasksMap[t.ID] = t.Name
//	}
//
//	for _, u := range entries {
//		dto := TimeEntryDTO{
//			Description: u.Description,
//			ID:          u.ID,
//			Project:     projectMap[u.Project],
//			ProjectId:   u.Project,
//			TaskId:      u.Task,
//			Task:        tasksMap[u.Task],
//			TimeSpent:   u.TimeSpent,
//		}
//
//		dtos = append(dtos, dto)
//	}
//
//	RespondWithJson(w, http.StatusOK, dtos)
//
//	return nil
//}

//func GetTimeEntryByIdEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
//	defer func() {
//		if r := recover(); r != nil {
//			RespondWithError(w, http.StatusNotFound, "Not Found")
//		}
//		if r.Body != nil{
//			r.Body.Close()
//		}
//	}()
//	params := mux.Vars(r)
//
//	entry, err := e.DBConnection.FindById(params["id"])
//
//	timeEntry := entry.(TimeEntry)
//
//	if err != nil {
//		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s", "Invalid entries")}
//	}
//
//	response, err := http.Get(string("http://project-service:3000/" + timeEntry.Project.Hex()))
//
//	if err != nil {
//		return StatusError{http.StatusFailedDependency, fmt.Errorf("%s", "Network error")}
//	}
//
//	project := Project{}
//	defer response.Body.Close()
//	if json.NewDecoder(response.Body).Decode(&project); err != nil {
//		return StatusError{http.StatusFailedDependency, fmt.Errorf("%s", "Network error")}
//	}
//
//	dto := TimeEntryDTO{
//		Description: timeEntry.Description,
//		ID:          timeEntry.ID,
//		ProjectId:   timeEntry.Project,
//		TaskId:      timeEntry.Task,
//		TimeSpent:   timeEntry.TimeSpent,
//	}
//
//	RespondWithJson(w, http.StatusOK, dto)
//
//	return nil
//}

func AddTimeEntryEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	var timeEntry TimeEntry
	if err := json.NewDecoder(r.Body).Decode(&timeEntry); err != nil {
		return StatusError{http.StatusBadRequest, fmt.Errorf("%s", "Invalid request payload, ")}
	}
	timeEntry.ID = bson.NewObjectId()

	if err := e.DBConnection.Insert(timeEntry); err != nil {
		return StatusError{http.StatusInternalServerError, fmt.Errorf("%s", err)}
	}
	RespondWithJson(w, http.StatusCreated, timeEntry)

	return nil
}

func UpdateTimeEntryEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()
	var entry TimeEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		return StatusError{http.StatusBadRequest, fmt.Errorf("%s", "Invalid request payload")}
	}

	if err := e.DBConnection.Update(entry); err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}

func DeleteTimeEntryEndPoint(e *Env, w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)

	if err := e.DBConnection.Delete(params["id"]); err != nil {
		return StatusError{http.StatusInternalServerError, err}
	}
	RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})

	return nil
}
