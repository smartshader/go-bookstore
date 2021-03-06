package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"smartshader/go-bookstore/users/domain/users"
	"smartshader/go-bookstore/users/services"
	"smartshader/go-bookstore/users/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 10)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	res, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, res.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	res, restErr := services.UsersService.GetUser(userId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, res.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr = errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	res, restErr := services.UsersService.UpdateUser(isPartial, user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, res.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	if restErr = services.UsersService.DeleteUser(userId); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
