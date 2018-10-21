package database

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"log"

	."time-logger/internal/pkg/entities"
)

type TimeEntryDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "time-entries"
)

func (m *TimeEntryDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TimeEntryDAO) FindAll() ([]TimeEntry, error) {
	var entries []TimeEntry
	err := db.C(COLLECTION).Find(bson.M{}).All(&entries)
	return entries, err
}

func (m* TimeEntryDAO) FindAllByProjectId(id string) ([]TimeEntry, error) {
	var entries []TimeEntry

	err := db.C(COLLECTION).Find(bson.M{"projectId": bson.ObjectIdHex(id)}).All(&entries)
	return entries, err
}

func (m *TimeEntryDAO) FindById(id string) (TimeEntry, error) {
	var entries TimeEntry
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entries)
	return entries, err
}

func (m *TimeEntryDAO) Insert(entry TimeEntry) error {
	err := db.C(COLLECTION).Insert(&entry)
	return err
}

func (m *TimeEntryDAO) Delete(entry TimeEntry) error {
	err := db.C(COLLECTION).Remove(&entry)
	return err
}

func (m *TimeEntryDAO) Update(entry TimeEntry) error {
	err := db.C(COLLECTION).UpdateId(entry.ID, &entry)
	return err
}