package database

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"testing"
	"time-logger/internal/pkg/database-access"
	"time-logger/internal/pkg/entities"
)

var taskDAO TaskDAO
var dataAccessLayerMock *database_access.DataLayerMock
var collectionMock *database_access.CollectionMock
var queryMock *database_access.QueryMock

func setupMany(results *[]interface{}, err error) {
	dataAccessLayerMock = &database_access.DataLayerMock{
		CFunc: func(name string) database_access.Collection {
			return collectionMock
		},
	}

	collectionMock = &database_access.CollectionMock{
		FindFunc: func(query interface{}) database_access.Query {
			return queryMock
		},
	}

	queryMock = &database_access.QueryMock{
		AllFunc: func(result interface{}) error {
			result = results
			return err
		},
	}

	taskDAO = TaskDAO{DB: dataAccessLayerMock}
}

func setupSingle(result interface{}, err error) {
	dataAccessLayerMock = &database_access.DataLayerMock{
		CFunc: func(name string) database_access.Collection {
			return collectionMock
		},
	}

	collectionMock = &database_access.CollectionMock{
		FindFunc: func(query interface{}) database_access.Query {
			return queryMock
		},
		RemoveFunc: func(selector interface{}) error {
			return err
		},
		UpdateIdFunc: func(id interface{}, update interface{}) error {
			return err
		},
		FindIdFunc: func(id interface{}) database_access.Query {
			return queryMock
		},
		UpdateFunc: func(selector interface{}, update interface{}) error {
			return err
		},
		InsertFunc: func(docs ...interface{}) error {
			return err
		},
	}

	queryMock = &database_access.QueryMock{
		AllFunc: func(result interface{}) error {
			result = result
			return err
		},
		OneFunc: func(result interface{}) error {
			return err
		},
	}

	taskDAO = TaskDAO{DB: dataAccessLayerMock}
}

func TestTaskDAO_FindAll(t *testing.T) {
	projects := make([] interface{}, 3)
	for i := range projects {
		projects[i] = entities.Task{Name: "Test" + strconv.Itoa(i)}
	}

	setupMany(&projects, nil)
	taskDAO.FindAll()

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 1, len(c))

	fc := collectionMock.FindCalls()
	assert.Equal(t, 1, len(fc))

	ac := queryMock.AllCalls()
	assert.Equal(t, 1, len(ac))
}

func TestTaskDAO_FindById_FindById(t *testing.T) {
	project := entities.Task{Name: "Test"}
	setupSingle(project, nil)

	taskDAO.FindById(bson.NewObjectId().Hex())

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 1, len(c))

	fc := collectionMock.FindIdCalls()
	assert.Equal(t, 1, len(fc))

	ac := queryMock.OneCalls()
	assert.Equal(t, 1, len(ac))
}
func TestTasksDAO_Insert(t *testing.T) {
	project := entities.Task{Name: "Test"}

	setupSingle(entities.Task{}, nil)

	taskDAO.Insert(project)

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 1, len(c))

	fc := collectionMock.InsertCalls()
	assert.Equal(t, 1, len(fc))
}

func TestTasksDAO_Insert_Error(t *testing.T) {
	te := entities.TimeEntry{}
	setupSingle(entities.Task{}, nil)

	err := taskDAO.Insert(te)

	assert.Equal(t, "Given type is not a entities.Task", err.Error())

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 0, len(c))

	fc := collectionMock.InsertCalls()
	assert.Equal(t, 0, len(fc))
}

func TestTasksDAO_Update(t *testing.T) {
	project := entities.Task{ID: bson.NewObjectId(), Name: "Test"}

	setupSingle(entities.Task{}, nil)

	taskDAO.Update(project)

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 1, len(c))

	fc := collectionMock.UpdateIdCalls()
	assert.Equal(t, 1, len(fc))
}

func TestTasksDAO_Update_Error(t *testing.T) {
	te := entities.TimeEntry{}
	setupSingle(entities.Task{}, nil)

	err := taskDAO.Update(te)

	assert.Equal(t, "Given type is not a entities.Task", err.Error())

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 0, len(c))

	fc := collectionMock.UpdateIdCalls()
	assert.Equal(t, 0, len(fc))
}

func TestTasksDAO_Delete(t *testing.T) {
	project := entities.Task{ID: bson.NewObjectId(), Name: "Test"}

	setupSingle(entities.Task{}, nil)

	taskDAO.Delete(project)

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 1, len(c))

	fc := collectionMock.RemoveCalls()
	assert.Equal(t, 1, len(fc))
}

func TestTasksDAO_Delete_Error(t *testing.T) {
	te := entities.TimeEntry{}
	setupSingle(entities.Task{}, nil)

	err := taskDAO.Delete(te)

	assert.Equal(t, "Given type is not a entities.Task", err.Error())

	c := dataAccessLayerMock.CCalls()
	assert.Equal(t, 0, len(c))

	fc := collectionMock.RemoveCalls()
	assert.Equal(t, 0, len(fc))
}
