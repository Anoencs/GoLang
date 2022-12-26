package models

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var userList = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
	{Username: "user3", Password: "pass3"},
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("THE PASSWORD CAN'T BE EMPTY")
	} else if !IsUsernameAvailable(username) {
		return nil, errors.New("THE USERNAME ISN'T AVAILABLE")
	}

	u := User{Username: username, Password: password}

	userList = append(userList, u)

	return &u, nil

}

func IsUsernameAvailable(username string) bool {
	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}
	return true

}

func Render(c *gin.Context, templates string, data gin.H) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templates, data)
	}
}
func IsUserValid(username, password string) bool {
	for _, user := range userList {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
