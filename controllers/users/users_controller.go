package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jgmc3012/bookstore_users-api/domain/users"
	"github.com/jgmc3012/bookstore_users-api/services"
	"github.com/jgmc3012/bookstore_users-api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
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
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	user.Id = userId

	result, getErr := services.GetUser(user)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
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
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := services.UpdateUser(isPartial, user)
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

	deleteErr := services.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func Search(c *gin.Context) {
	status := c.Query("status")

	results, err := services.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, results.Marshal(c.GetHeader("X-Public") == "true"))
}
