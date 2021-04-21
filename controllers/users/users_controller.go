package users

import (
	"net/http"

	"github.com/aushafy/go-bookstore-users-api/domain/users"
	"github.com/aushafy/go-bookstore-users-api/services"
	"github.com/aushafy/go-bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context) {
	// logging with logrush with JSON formatter
	log.SetFormatter(&log.JSONFormatter{})

	var user users.User

	// // catch all user HTTP body request
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	log.Error(err)
	// }

	// // then unmarshal JSON body request to the User Struct
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	log.Error(err)
	// }

	// alternative way to simplify above code, catch all user request as a JSON format without using ioutil.ReadAll and unmarshal it
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.RestErr{
			Message: "Invalid JSON Body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
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
