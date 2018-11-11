package database_access

import (
	"gopkg.in/mgo.v2"
)

// Session is an interface to access to the Session struct.
//go:generate moq -out session_mongo_mock.go . Session
type Session interface {
	DB(name string) DataLayer
	Close()
}

func NewSession(server string) (MongoSession, error){
	s,e := mgo.Dial(server)
	return MongoSession{ s}, e
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (s MongoSession) DB(name string) DataLayer {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// DataLayer is an interface to access to the database struct.
//go:generate moq -out datalayer_mongo_mock.go . DataLayer
type DataLayer interface {
	C(name string) Collection
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// MongoQuery wraps a mgo.Query to embed methods in models.
type MongoQuery struct {
	*mgo.Query
}

// Collection is an interface to access to the collection struct.
//go:generate moq -out collection_mongo_mock.go . Collection
type Collection interface {
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
	UpdateId(id interface{}, update interface{}) error
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a DataLayer interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// Find shadows *mgo.Collection to returns a Query interface instead of *mgo.Query.
func (c MongoCollection) Find(query interface{}) Query {
	return MongoQuery{Query: c.Collection.Find(query)}
}

func (c MongoCollection) FindId(query interface{}) Query {
	return MongoQuery{Query: c.Collection.FindId(query)}
}


// Query is an interface to access to the database struct
//go:generate moq -out query_mongo_mock.go . Query
type Query interface {
	All(result interface{}) error
	One(result interface{}) (err error)
	Distinct(key string, result interface{}) error
}
