package entities

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Project struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	StartDate   time.Time     `bson:"startDate" json:"startDate"`
	EndDate     time.Time     `bson:"endDate" json:"endDate"`
	Budget      float64       `bson:"bugdet" json:"budget,string"`
	Price       float64       `bson:"price" json:"price,string"`
	Stakeholder string        `bson:"stakeholder" json:"stakeholder"`
}
