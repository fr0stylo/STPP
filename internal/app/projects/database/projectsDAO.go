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

func (m *ProjectDAO) FindAll() (interface{}, error) {
	var entries []Project
	err := m.DB.C(COLLECTION).Find(bson.M{}).All(&entries)
	if err != nil {
		return nil, fmt.Errorf("Unable to get all projects, %s", err)
	}

	return &entries, nil
}

func (m *ProjectDAO) FindById(id string) (interface{}, error) {
	var entries Project
	err := m.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)

	return &entries, err
}

func (m *ProjectDAO) Insert(entry interface{}) error {
	project, ok := entry.(Project)
	if !ok {
		return fmt.Errorf("%s", "Given type is not a project")
	}

	err := m.DB.C(COLLECTION).Insert(&project)
	return err
}

func (m *ProjectDAO) Delete(id string) error {
	var objectId bson.ObjectId

	if objectId = bson.ObjectIdHex(id); recover() != nil {
		return fmt.Errorf("%s", "Cannot convert string to object id")
	}

	err := m.DB.C(COLLECTION).Remove(bson.M{"_id": objectId})
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
