package user

import (
	"context"
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/ypeckstadt/golang-graphql-study/pkg/models"
	"net/http"
)

// Custom errors
var (
	ErrUserNotFound = echo.NewHTTPError(http.StatusForbidden, "User not found.")
)

func (u *User) Create(name string) (models.User, error) {
	user := models.User{
		UUID: u.uuid.Generate(),
		Name: name,
	}
	err := user.Insert(context.Background(), u.db, boil.Infer())
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *User) Update(uuid string, request UpdateRequest) (models.User, error) {
	// get user from database
	user, err := models.Users(qm.Where(models.UserColumns.UUID+"=?", uuid)).One(context.Background(), u.db)
	if err != nil && err == sql.ErrNoRows {
		return models.User{}, ErrUserNotFound
	}

	// set updated fields
	user.Name = request.Name

	// update in database
	_, err = user.Update(context.Background(), u.db, boil.Infer())

	return *user, err
}

func (u *User) Get(uuid string) (models.User, error) {
	// get user from database
	user, err := models.Users(qm.Where(models.UserColumns.UUID+"=?", uuid)).One(context.Background(), u.db)
	if err != nil && err == sql.ErrNoRows {
		return models.User{}, ErrUserNotFound
	}

	return *user, nil
}

// List returns all users in the database
func (u *User) List() (models.UserSlice, error) {
	return models.Users().All(context.Background(), u.db)
}

func (u *User) Delete(uuid string) error {

	// get user
	user, err := models.Users(qm.Where(models.UserColumns.UUID+"=?", uuid)).One(context.Background(), u.db)
	if err != nil && err == sql.ErrNoRows {
		return ErrUserNotFound
	}

	// delete
	_, err = user.Delete(context.Background(), u.db)

	return err
}
