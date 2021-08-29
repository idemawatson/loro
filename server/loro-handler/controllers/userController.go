package controllers

import (
	"loro-handler/models"
	"loro-handler/repository"
	"loro-handler/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us service.UserService
}

type GetUserInfoRequest struct {
	UserId string `json:"userId"`
}

func NewUserController(userRepo repository.UserRepo) *UserController {
	us := service.NewUserService(userRepo)
	return &UserController{us}
}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	var req models.GetUserInfoRequest
	c.ShouldBindJSON(&req)
	userInfo, err := uc.us.GetUserInfo(&req)
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
