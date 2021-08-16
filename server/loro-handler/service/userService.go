package service

import (
	"log"
	"loro-handler/database"
	"loro-handler/models"
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

func (UserService) ListUser() ([]models.User, error) {
	table := database.GetMainTable()
	var results []models.User
	err := table.Get("id", "000001").All(&results)
	return results, err
}
