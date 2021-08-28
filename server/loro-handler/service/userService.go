package service

import (
	"fmt"
	"loro-handler/models"
	"loro-handler/repository"
)

type UserService interface {
	GetUserInfo(req *models.GetUserInfoRequest) (models.User, error)
}

type userService struct {
	ur repository.UserRepo
}

func NewUserService(ur repository.UserRepo) UserService {
	return &userService{ur}
}

func (us *userService) GetUserInfo(req *models.GetUserInfoRequest) (models.User, error) {
	fmt.Printf("[INFO] GetUserInfo request: %v\n", req.UserId)
	result, err := us.ur.GetUserInfo(req.UserId)
	fmt.Printf("[INFO] GetUserInfo response: %v\n", result)
	return result, err
}
