package users

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aushafy/go-bookstore-users-api/domain/users"
	"github.com/aushafy/go-bookstore-users-api/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func CreateUser(c *gin.Context) {
	var user users.User
	// catch all user HTTP body request
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error(err)
	}
	// then unmarshal JSON body request to the User Struct
	if err := json.Unmarshal(bytes, &user); err != nil {
		log.Error(err)
	}
	result, err := services.CreateUser(user)
	if err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusNotImplemented, "Implement me!")
// }
