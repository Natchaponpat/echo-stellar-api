package user

import (
	"net/http"

	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/storage"

	"github.com/Natchaponpat/echo-stellar-api/echo/better-echo/example/model"

	"github.com/labstack/echo"
)

type Handler struct {
	userStorage storage.IUserStorage
}

func New(storage storage.IUserStorage) (*Handler, error) {
	return &Handler{
		userStorage: storage,
	}, nil
}

func (h *Handler) Handle(e *echo.Echo) {
	e.GET("/users", h.list())
	e.POST("/users", h.add())
	e.GET("/users/:name", h.get())
}

func (h *Handler) list() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.userStorage.List()
		if err != nil {
			return err
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

		err = h.userStorage.Create(user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		res, err := h.userStorage.Get(name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, res)
	}
}
