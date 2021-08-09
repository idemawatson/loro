package service

import (
	"log"
	"loro-handler/database"
	"loro-handler/models"
)

type UserService struct{}

// func NewUser(userId string, userName string, description string, mailAddress string, profilePicturePath string) *models.User {
// 	u := new(models.User)
// 	u.Id = "111"
// 	u.UserId = userId
// 	u.UserName = userName
// 	u.Description = description
// 	u.MailAddress = mailAddress
// 	u.ProfilePicturePath = profilePicturePath
// 	u.NumPost = 0
// 	u.NumFollowing = 0
// 	u.NumFollowed = 0
// 	u.NumFavorite = 0
// 	return u
// }

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
