package service

import (
	"fmt"
	"log"
	"loro-handler/database"
	"loro-handler/models"

	"github.com/guregu/dynamo"
)

type UserService struct{}

func (UserService) CreateUser() error {
	user := models.User{
		Id:     "000001",
		UserId: "000001",
	}
	table := database.GetMainTable()
	log.Printf("user: %+v\n", user)
	err := table.Put(user).Run()
	if err != nil {
		return err
	}
	return nil
}

func (UserService) GetUserInfo(req *models.GetUserInfoRequest) (models.User, error) {
	table := database.GetMainTable()
	var result models.User
	fmt.Printf("[INFO] GetUserInfo request: %v\n", req.UserId)
	err := table.Get("userId", req.UserId).Range("profile", dynamo.Equal, "true").Index("userId-profile-Index").One(&result)
	return result, err
}
