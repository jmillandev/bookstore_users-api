package users

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jgmc3012/bookstore_users-api/domain/users"
	"github.com/jgmc3012/bookstore_users-api/services"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//Handle json error
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
