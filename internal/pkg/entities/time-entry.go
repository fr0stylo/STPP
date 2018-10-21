package entities

import "gopkg.in/mgo.v2/bson"

type TimeEntry struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Project bson.ObjectId `bson:"projectId" json:"projectId"`
	Task bson.ObjectId `bson:"taskid" json:"taskId"`
	TimeSpent bson.Decimal128 `bson:"timeSpent" json:"timeSpent"`
	Description string `bson:"description" json:"description"`
}
