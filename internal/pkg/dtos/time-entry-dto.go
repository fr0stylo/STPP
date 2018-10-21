package dtos

import "gopkg.in/mgo.v2/bson"

type TimeEntryDTO struct{
	ID bson.ObjectId `bson:"_id" json:"id"`
	ProjectId bson.ObjectId `bson:"projectId" json:"projectId"`
	Project string `bson:"project" json:"project"`
	TaskId bson.ObjectId `bson:"taskId" json:"taskId"`
	Task string `bson:"task" json:"task"`
	TimeSpent bson.Decimal128 `bson:"timeSpent" json:"timeSpent"`
	Description string `bson:"description" json:"description"`
}
