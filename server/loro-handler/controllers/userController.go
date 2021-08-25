package controllers

import (
	"fmt"
	"loro-handler/models"
	"loro-handler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}
type GetUserInfoRequest struct {
	UserId string `json:"userId"`
}

func NewUserController() *UserController {
	return new(UserController)
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	var req models.GetUserInfoRequest
	c.ShouldBindJSON(&req)
	fmt.Println(req)
	userService := service.UserService{}
	userInfo, err := userService.GetUserInfo(&req)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		userInfo,
	))
}
