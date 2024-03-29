package user

import (
	"net/http"

	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/model"

	"github.com/globalsign/mgo"

	"github.com/Natchaponpat/echo-stellar-api/echo/mongo/storage"

	"github.com/labstack/echo"
)

type Handler struct {
	userStorage storage.IUserStorage
}

func New(userStorage storage.IUserStorage) (*Handler, error) {
	return &Handler{userStorage}, nil
}

func (h *Handler) Handle(e *echo.Echo) {
	e.GET("/users", h.list())
	e.POST("/users", h.add())
	e.GET("/users/:name", h.get())
}

func (h *Handler) list() echo.HandlerFunc {
	return func(c echo.Context) error {
		list, err := h.userStorage.List()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		if list == nil {
			list = []model.User{}
		}
		return c.JSON(http.StatusOK, list)
	}
}

func (h *Handler) add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		err := c.Bind(&user)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		u, err := h.userStorage.Get(user.Name)
		if err != nil {
			if err != mgo.ErrNotFound {
				c.Logger().Error(err)
				return echo.NewHTTPError(http.StatusInternalServerError)
			}
		}
		if u.Name != "" {
			return echo.NewHTTPError(http.StatusBadRequest, "name has been used")
		}

		err = h.userStorage.Create(user)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		user, err := h.userStorage.Get(name)
		if err != nil {
			if err == mgo.ErrNotFound {
				return echo.NewHTTPError(http.StatusNotFound, "user not found")
			}
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
		return c.JSON(http.StatusOK, user)

	}
}
