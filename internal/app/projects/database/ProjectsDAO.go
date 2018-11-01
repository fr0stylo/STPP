package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"

	. "time-logger/internal/pkg/entities"
)

type IProjectDAO interface {
	Connect()
	FindAll() ([] Project, error)
	FindById(id string) (Project, error)
	Insert(entry Project) error
	Update(entry Project) error
	Delete(entry Project) error
}

type ProjectDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "project"
)

func (m *ProjectDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *ProjectDAO) FindAll() ([]Project, error) {
	var entries []Project
	err := db.C(COLLECTION).Find(bson.M{}).All(&entries)
	return entries, err
}

func (m *ProjectDAO) FindById(id string) (Project, error) {
	var entries Project
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)
	return entries, err
}

func (m *ProjectDAO) Insert(entry Project) error {
	err := db.C(COLLECTION).Insert(&entry)
	return err
}

func (m *ProjectDAO) Delete(entry Project) error {
	err := db.C(COLLECTION).Remove(&entry)
	return err
}

func (m *ProjectDAO) Update(entry Project) error {
	err := db.C(COLLECTION).UpdateId(entry.ID, &entry)
	return err
}
