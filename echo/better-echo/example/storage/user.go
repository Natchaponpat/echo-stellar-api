package storage

import (
	"net/http"

	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/model"
	"github.com/labstack/echo"
)

type IUserStorage interface {
	List() ([]model.User, error)
	Get(name string) (model.User, error)
	Create(model.User) error
}

type UserStorage struct {
	listUser []model.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		listUser: []model.User{},
	}
}

func (s *UserStorage) List() ([]model.User, error) {
	return s.listUser, nil
}

func (s *UserStorage) Get(name string) (model.User, error) {
	var user model.User
	var found bool
	for _, u := range s.listUser {
		if u.Name == name {
			user = u
			found = true
			break
		}
	}
	if !found {
		return model.User{}, echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return user, nil
}

func (s *UserStorage) Create(user model.User) error {
	for _, u := range s.listUser {
		if user.Name == u.Name {
			return echo.NewHTTPError(http.StatusBadRequest, "name has been used")
		}
	}

	s.listUser = append(s.listUser, user)
	return nil
}
