package user

import (
	"net/http"

	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/model"

	"github.com/labstack/echo"
)

type Handler struct {
	listUser []model.User
}

func New() (*Handler, error) {
	return &Handler{
		listUser: []model.User{},
	}, nil
}

func (h *Handler) Handle(e *echo.Echo) {
	e.GET("/users", h.list())
	e.POST("/users", h.add())
	e.GET("/users/:name", h.get())
}

func (h *Handler) list() echo.HandlerFunc {
	return func(c echo.Context) error {
		res := h.listUser
		if res == nil {
			res = []model.User{}
		}
		return c.JSON(http.StatusOK, res)
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

		for _, u := range h.listUser {
			if user.Name == u.Name {
				return echo.NewHTTPError(http.StatusBadRequest, "name has been used")
			}
		}

		h.listUser = append(h.listUser, user)
		return c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		var user model.User
		var found bool
		for _, u := range h.listUser {
			if u.Name == name {
				user = u
				found = true
				break
			}
		}
		if !found {
			return echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return c.JSON(http.StatusOK, user)

	}
}
