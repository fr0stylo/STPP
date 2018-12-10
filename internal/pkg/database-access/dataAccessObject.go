package database_access

//go:generate moq -out dataAccessObject_mock.go . DataAccessObject
type DataAccessObject interface {
	FindAll() (interface{}, error)
	FindById(id string) (interface{}, error)
	Insert(entry interface{}) error
	Update(entry interface{}) error
	Delete(id string) error
}
