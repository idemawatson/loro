package service

import (
	"loro-handler/models"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type UserRepoMock struct{}

func (u UserRepoMock) GetUserInfo(userId string) (models.User, error) {
	return models.User{UserId: userId}, nil
}

func TestUserService_GetUserInfo(t *testing.T) {
	//setup
	us := &userService{ur: UserRepoMock{}}
	userId := "000001"
	request := models.GetUserInfoRequest{UserId: userId}
	expected := models.User{
		UserId: userId,
	}

	//when
	res, err := us.GetUserInfo(&request)

	//then
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(res, expected); diff != "" {
		t.Errorf("Diff: (-got +want)\n%s", diff)
	}
}

func TestNewUserService_userServiceが取得できること(t *testing.T) {
	//when
	res := NewUserService(UserRepoMock{})
	expected := &userService{ur: UserRepoMock{}}

	//then
	if diff := cmp.Diff(res, expected, cmpopts.IgnoreUnexported(userService{})); diff != "" {
		t.Errorf("Diff: (-got +want)\n%s", diff)
	}
}
