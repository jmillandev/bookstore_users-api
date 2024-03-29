package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmillandev/bookstore_oauth-go/oauth"
	"github.com/jmillandev/bookstore_users-api/domain/users"
	"github.com/jmillandev/bookstore_users-api/services"
	"github.com/jmillandev/bookstore_utils-go/rest_errors"
)

func getUserId(userIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Create(c *gin.Context) {
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
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}

	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := rest_errors.NewUnauthorizedError()
		c.JSON(err.Status, err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	user.Id = userId

	result, getErr := services.UsersService.GetUser(user)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == result.Id {
		c.JSON(http.StatusOK, result.Marshal(false))
		return
	}

	c.JSON(http.StatusOK, result.Marshal(oauth.IsPublic(c.Request)))
}

func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println(err)
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	deleteErr := services.UsersService.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func Search(c *gin.Context) {
	status := c.Query("status")

	results, err := services.UsersService.FindByStatusUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, results.Marshal(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}
