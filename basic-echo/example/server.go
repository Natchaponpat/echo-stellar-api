package main

// ==========================================================
// This is exercise for learning about basic of echo web framework
// This will consist of 3 apis about managing User.
// 1. List users
// 2. Add new user to the list
// 3. Get user's info by name
// ==========================================================

import (
	"net/http"

	"github.com/labstack/echo"
)

// Create new type struct called User
// which consist of two variable: Name as string and Age as integer
// Also included json tag.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Init empty List for store new User.
var listUser []User

func main() {
	// Init new Echo instance
	e := echo.New()

	// Create new route using method GET to list the existing users
	// and return slice of users in json format as a response.
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, listUser)
	})

	// Create new route using method POST to receive name and age for creating user.
	// and add new user to the list.Then return the new user as a response.
	e.POST("/users", func(c echo.Context) error {
		var user User
		err := c.Bind(&user)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		for _, u := range listUser {
			if user.Name == u.Name {
				return echo.NewHTTPError(http.StatusBadRequest, "name has been used")
			}
		}

		listUser = append(listUser, user)
		return c.JSON(http.StatusOK, user)
	})

	// Create new route using method GET to search user in list by name
	// and return user data as a response.
	e.GET("/users/:name", func(c echo.Context) error {
		name := c.Param("name")

		for _, u := range listUser {
			if u.Name == name {
				return c.JSON(http.StatusOK, u)
			}
		}

		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	})

	// Start new server
	e.Logger.Fatal(e.Start(":8000"))
}
