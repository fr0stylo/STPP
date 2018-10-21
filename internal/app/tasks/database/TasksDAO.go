package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"

	. "time-logger/internal/pkg/entities"
)

type TaskDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "task"
)

func (m *TaskDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TaskDAO) FindAll() ([]Task, error) {
	var entries []Task
	err := db.C(COLLECTION).Find(bson.M{}).All(&entries)
	return entries, err
}

func (m *TaskDAO) FindById(id string) (Task, error) {
	var entries Task
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)
	return entries, err
}

func (m *TaskDAO) Insert(entry Task) error {
	err := db.C(COLLECTION).Insert(&entry)
	return err
}

func (m *TaskDAO) Delete(entry Task) error {
	err := db.C(COLLECTION).Remove(&entry)
	return err
}

func (m *TaskDAO) Update(entry Task) error {
	err := db.C(COLLECTION).UpdateId(entry.ID, &entry)
	return err
}