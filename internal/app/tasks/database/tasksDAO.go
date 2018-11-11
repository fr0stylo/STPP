package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time-logger/internal/pkg/database-access"

	. "time-logger/internal/pkg/entities"
)

type TaskDAO struct {
	DB database_access.DataLayer
}

const (
	COLLECTION = "task"
)

func (m *TaskDAO) FindAll() ([] interface{}, error) {
	var entries []Task
	err := m.DB.C(COLLECTION).Find(bson.M{}).All(&entries)

	results := make([] interface{}, len(entries))

	for i, o := range entries {
		results[i] = o
	}
	return results, err
}

func (m *TaskDAO) FindById(id string) (interface{}, error) {
	var entries Task
	err := m.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)
	return entries, err
}

func (m *TaskDAO) Insert(entry interface{}) error {
	task, ok := entry.(Task)

	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", Task{})
	}

	err := m.DB.C(COLLECTION).Insert(&task)
	return err
}

func (m *TaskDAO) Delete(entry interface{}) error {
	task, ok := entry.(Task)

	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", Task{})
	}

	err := m.DB.C(COLLECTION).Remove(&task)
	return err
}

func (m *TaskDAO) Update(entry interface{}) error {
	task, ok := entry.(Task)

	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", Task{})
	}

	err := m.DB.C(COLLECTION).UpdateId(task.ID, &task)
	return err
}
