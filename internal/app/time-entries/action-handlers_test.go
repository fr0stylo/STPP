package time_entries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time-logger/internal/pkg/database-access"
	"time-logger/internal/pkg/entities"
	"time-logger/internal/pkg/http-wrappers"
)

var mockedIDataAccessObject *database_access.DataAccessObjectMock

func prepEnv(projects [] entities.TimeEntry, e error) *http_wrappers.Env {
	mockedIDataAccessObject = &database_access.DataAccessObjectMock{
		FindAllFunc: func() ([]interface{}, error) {
			result := make([] interface{}, len(projects))

			for i, o := range projects {
				result[i] = o
			}

			return result, e
		},
	}

	env := &http_wrappers.Env{DBConnection: mockedIDataAccessObject}

	return env
}

func prepEnvSingle(project entities.TimeEntry, e error) *http_wrappers.Env {
	mockedIDataAccessObject = &database_access.DataAccessObjectMock{
		FindByIdFunc: func(id string) (interface{}, error) {
			return project, e
		},
		InsertFunc: func(entry interface{}) error {
			return e
		},
		DeleteFunc: func(entry interface{}) error {
			return e
		},
		UpdateFunc: func(entry interface{}) error {
			return e
		},
	}

	env := &http_wrappers.Env{DBConnection: mockedIDataAccessObject}

	return env
}

//
//func TestGetAllTasksEndPoint_ShouldReturnStatusOk(t *testing.T) {
//	req, err := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	resp := httptest.NewRecorder()
//
//	expectedProjects := make([] entities.TimeEntry, 5)
//	for i := range expectedProjects {
//		expectedProjects[i] = entities.TimeEntry{Description: "Test" + strconv.Itoa(i)}
//	}
//
//	env := prepEnv(expectedProjects, nil)
//
//	handler := http_wrappers.Handler{env, GetAllTimeEntriesEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	assert.Equal(t, 200, resp.Code, "Expected 200")
//}
//
//func TestGetAllTasksEndPoint_ShouldReturnCorrectResutls(t *testing.T) {
//	expectedProjects := make([] entities.TimeEntry, 5)
//	for i := range expectedProjects {
//		expectedProjects[i] = entities.TimeEntry{Description: "Test" + strconv.Itoa(i)}
//	}
//
//	env := prepEnv(expectedProjects, nil)
//	req, err := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	resp := httptest.NewRecorder()
//
//	handler := http_wrappers.Handler{env, GetAllTimeEntriesEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	items := [] entities.TimeEntry{}
//	err = json.NewDecoder(resp.Body).Decode(&items)
//
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	assert.ElementsMatch(t, items, expectedProjects)
//}
//
//func TestGetAllTasksEndPoint_DeferShouldCloseBody(t *testing.T) {
//	expectedProjects := make([] entities.TimeEntry, 5)
//	for i := range expectedProjects {
//		expectedProjects[i] = entities.TimeEntry{Description: "Test" + strconv.Itoa(i)}
//	}
//
//	env := prepEnv(expectedProjects, nil)
//	req, err := http.NewRequest(http.MethodGet, "localhost:3000", ioutil.NopCloser(strings.NewReader("a")))
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	resp := httptest.NewRecorder()
//
//	handler := http_wrappers.Handler{env, GetAllTimeEntriesEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	items := [] entities.TimeEntry{}
//	err = json.NewDecoder(resp.Body).Decode(&items)
//
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	assert.ElementsMatch(t, items, expectedProjects)
//}
//
//func TestGetAllTasksEndPoint_ShouldReturnError(t *testing.T) {
//	expectedProjects := make([] entities.TimeEntry, 0)
//
//	env := prepEnv(expectedProjects, fmt.Errorf("%s", "No entries"))
//	req, _ := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//
//	resp := httptest.NewRecorder()
//
//	handler := http_wrappers.Handler{env, GetAllTimeEntriesEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	assert.JSONEq(t, "{\"error\":\"Invalid entries\"}", resp.Body.String())
//}
//
//func TestGetProjectEndPoint(t *testing.T) {
//	req, err := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	resp := httptest.NewRecorder()
//
//	project := entities.TimeEntry{Description: "TestProject"}
//
//	env := prepEnvSingle(project, nil)
//
//	handler := http_wrappers.Handler{env, GetTimeEntryByIdEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	assert.Equal(t, 200, resp.Code, "Expected 200")
//}
//
//func TestGetTasksEndPoint_ShouldReturnCorrectResutls(t *testing.T) {
//	expectedProject := entities.TimeEntry{Description: "Project Test"}
//
//	env := prepEnvSingle(expectedProject, nil)
//	req, err := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	resp := httptest.NewRecorder()
//
//	handler := http_wrappers.Handler{env, GetTimeEntryByIdEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	item := entities.TimeEntry{}
//	err = json.NewDecoder(resp.Body).Decode(&item)
//
//	if err != nil {
//		assert.Error(t, err)
//	}
//
//	assert.Equal(t, expectedProject, item)
//}
//
//func TestGetProjectEndPoint_ShouldReturnError(t *testing.T) {
//	env := prepEnvSingle(entities.TimeEntry{}, fmt.Errorf("%s", "No entries"))
//	req, _ := http.NewRequest(http.MethodGet, "localhost:3000", nil)
//
//	resp := httptest.NewRecorder()
//
//	handler := http_wrappers.Handler{env, GetTimeEntryByIdEndPoint}
//	handler.ServeHTTP(resp, req)
//
//	assert.JSONEq(t, "{\"error\":\"Invalid entries\"}", resp.Body.String())
//}

func TestAddProjectEndPoint(t *testing.T) {
	newProject := entities.TimeEntry{Description: "name"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, AddTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	d := mockedIDataAccessObject.InsertCalls()

	assert.Equal(t, 1, len(d))
}

func TestAddProjectEndPoint_ShouldFailWithErrorWhileInserting(t *testing.T) {
	newProject := entities.TimeEntry{Description: "Description"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, fmt.Errorf("%s", "Error Occurred"))

	handler := http_wrappers.Handler{env, AddTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.JSONEq(t, "{\"error\":\"Error Occurred\"}", resp.Body.String())

	d := mockedIDataAccessObject.InsertCalls()

	assert.Equal(t, 1, len(d))
}

func TestAddProjectEndPoint_ShouldFailWithIncorrectPayload(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(strings.NewReader("Name:\"name\",;Budget:1.0, Price:1000,Stakeholder:\"stake\"}")))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, AddTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	d := mockedIDataAccessObject.InsertCalls()
	assert.Equal(t, 0, len(d))
	assert.JSONEq(t, "{\"error\":\"Invalid request payload, \"}", resp.Body.String())
}

func TestUpdateProjectEndPoint(t *testing.T) {
	newProject := entities.TimeEntry{Description: "name"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, UpdateTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	d := mockedIDataAccessObject.UpdateCalls()

	assert.Equal(t, 1, len(d))
}

func TestUpdateProjectEndPoint_ShouldFailWithErrorWhileUpdating(t *testing.T) {
	newProject := entities.TimeEntry{Description: "name"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, fmt.Errorf("%s", "Error Occurred"))

	handler := http_wrappers.Handler{env, UpdateTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.JSONEq(t, "{\"error\":\"Error Occurred\"}", resp.Body.String())

	d := mockedIDataAccessObject.UpdateCalls()

	assert.Equal(t, 1, len(d))
}

func TestUpdateProjectEndPoint_ShouldFailWithIncorrectPayload(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(strings.NewReader("Name:\"name\",;Budget:1.0, Price:1000,Stakeholder:\"stake\"}")))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, UpdateTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	d := mockedIDataAccessObject.UpdateCalls()
	assert.Equal(t, 0, len(d))
	assert.JSONEq(t, "{\"error\":\"Invalid request payload\"}", resp.Body.String())
}

func TestDeleteProjectEndPoint(t *testing.T) {
	newProject := entities.TimeEntry{Description: "name"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, DeleteTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	d := mockedIDataAccessObject.DeleteCalls()

	assert.Equal(t, 1, len(d))
}

func TestDeleteProjectEndPoint_ShouldFailWithErrorWhileUpdating(t *testing.T) {
	newProject := entities.TimeEntry{Description: "name"}
	payload, err := json.Marshal(newProject)
	if err != nil {
		assert.Error(t, err)
	}
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(bytes.NewReader(payload)))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, fmt.Errorf("%s", "Error Occurred"))

	handler := http_wrappers.Handler{env, DeleteTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	assert.JSONEq(t, "{\"error\":\"Error Occurred\"}", resp.Body.String())

	d := mockedIDataAccessObject.DeleteCalls()

	assert.Equal(t, 1, len(d))
}

func TestDeleteProjectEndPoint_ShouldFailWithIncorrectPayload(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "localhost:3000", ioutil.NopCloser(strings.NewReader("Name:\"name\",;Budget:1.0, Price:1000,Stakeholder:\"stake\"}")))

	resp := httptest.NewRecorder()

	env := prepEnvSingle(entities.TimeEntry{}, nil)

	handler := http_wrappers.Handler{env, DeleteTimeEntryEndPoint}
	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	d := mockedIDataAccessObject.UpdateCalls()
	assert.Equal(t, 0, len(d))
	assert.JSONEq(t, "{\"error\":\"Invalid request payload\"}", resp.Body.String())
}
