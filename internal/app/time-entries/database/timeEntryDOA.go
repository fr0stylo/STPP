package database

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time-logger/internal/pkg/database-access"

	. "time-logger/internal/pkg/entities"
)

type TimeEntryDAO struct {
	DB database_access.DataLayer
}

const (
	COLLECTION = "time-entries"
)

func (m *TimeEntryDAO) FindAll() ([]interface{}, error) {
	var entries []TimeEntry
	err := m.DB.C(COLLECTION).Find(bson.M{}).All(&entries)

	results := make([] interface{}, len(entries))

	for i, o := range entries {
		results[i] = o
	}

	return results, err
}

func (m *TimeEntryDAO) FindById(id string) (interface{}, error) {
	var entry TimeEntry
	err := m.DB.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entry)

	return entry, err
}

func (m *TimeEntryDAO) Insert(entry interface{}) error {
	timeEntry, ok := entry.(TimeEntry)
	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", TimeEntry{})
	}

	err := m.DB.C(COLLECTION).Insert(&timeEntry)

	return err
}

func (m *TimeEntryDAO) Delete(entry interface{}) error {
	timeEntry, ok := entry.(TimeEntry)
	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", TimeEntry{})
	}

	err := m.DB.C(COLLECTION).Remove(&timeEntry)
	return err
}

func (m *TimeEntryDAO) Update(entry interface{}) error {
	timeEntry, ok := entry.(TimeEntry)
	if !ok {
		return fmt.Errorf("%s %T", "Given type is not a", TimeEntry{})
	}

	err := m.DB.C(COLLECTION).UpdateId(timeEntry.ID, &entry)
	return err
}
