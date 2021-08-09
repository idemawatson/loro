package controllers

import (
	"loro-handler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return new(UserController)
}

func (uc *UserController) ListUser(c *gin.Context) {
	userService := service.UserService{}
	userList, err := userService.ListUser()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		userList,
	))
}
