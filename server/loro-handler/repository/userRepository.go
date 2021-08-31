package repository

import (
	"loro-handler/models"

	"github.com/guregu/dynamo"
)

type UserRepo interface {
	GetUserInfo(string) (models.User, error)
}

type userRepo struct {
	table dynamoTableClient
}

func NewUserRepository(table *dynamoTableClient) UserRepo {
	return &userRepo{*table}
}

func (ur *userRepo) GetUserInfo(userId string) (models.User, error) {
	var result models.User
	err := ur.table.Get("userId", userId).Range("profile", dynamo.Equal, "true").Index("userId-profile-Index").One(&result)
	return result, err
}
