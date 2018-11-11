package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time-logger/internal/pkg/database-access"

	. "time-logger/internal/pkg/entities"
)

type ProjectDAO struct {
	DB database_access.DataLayer
}

const (
	COLLECTION = "project"
)

func (m *ProjectDAO) FindAll() ([] interface{}, error) {
	var entries []Project
	err := m.DB.C(COLLECTION).Find(bson.M{}).All(&entries)

	result := make([] interface{}, len(entries))

	for i, o := range result {
		result[i] = o
	}

	return result, err
}

func (m *ProjectDAO) FindById(id string) (interface{}, error) {
	var entries Project
	err := m.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)
	return entries, err
}

func (m *ProjectDAO) Insert(entry interface{}) error {
	project, ok := entry.(Project)
	if !ok {
		return fmt.Errorf("%s", "Given type is not a project")
	}

	err := m.DB.C(COLLECTION).Insert(&project)
	return err
}

func (m *ProjectDAO) Delete(entry interface{}) error {
	project, ok := entry.(Project)
	if !ok {
		return fmt.Errorf("%s", "Given type is not a project")
	}

	err := m.DB.C(COLLECTION).Remove(&project)
	return err
}

func (m *ProjectDAO) Update(entry interface{}) error {
	project, ok := entry.(Project)
	if !ok {
		return fmt.Errorf("%s", "Given type is not a project")
	}

	err := m.DB.C(COLLECTION).UpdateId(project.ID, &project)
	return err
}
