package user

import (
	"database/sql"
	"github.com/ypeckstadt/golang-graphql-study/pkg/boiled"
)

// New creates new User service
func New(db *sql.DB, uuid UUID) User {
	return User{
		db: db,
		uuid: uuid,
	}
}

// Initialize initializes user application service
func Initialize(db *sql.DB, uuid UUID) User {
	return New(db, uuid)
}

// Service represents user service interface
type Service interface {
	Create(name string) (boiled.User, error)
	Update(uuid string, request UpdateRequest) (boiled.User, error)
	Get(uuid string) (boiled.User, error)
	List() (boiled.UserSlice, error)
	Delete(uuid string) error
}

// User represents user application service
type User struct {
	db   *sql.DB
	uuid UUID
}

// UUID represents the generator for UUIDs service interface
type UUID interface {
	Generate() string
}

type UpdateRequest struct {
	Name string
}
