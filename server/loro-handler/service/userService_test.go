package service

import (
	mock_repository "loro-handler/mock/repository"
	"loro-handler/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestUserService(t *testing.T) {
	t.Run("TestUserService_GetUserInfo_ユーザ情報が取得できること", func(t *testing.T) {
		//setup
		ctrl := gomock.NewController(t)

		mUr := mock_repository.NewMockUserRepo(ctrl)
		us := &userService{ur: mUr}

		userId := "000001"
		request := models.GetUserInfoRequest{UserId: userId}
		expected := models.User{
			UserId: userId,
		}
		mUr.EXPECT().GetUserInfo(userId).Return(expected, nil)

		//when
		res, err := us.GetUserInfo(&request)

		//then
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(res, expected); diff != "" {
			t.Errorf("Diff: (-got +want)\n%s", diff)
		}
	})
	t.Run("TestNewUserService_userServiceが取得できること", func(t *testing.T) {
		//setup
		ctrl := gomock.NewController(t)
		mUr := mock_repository.NewMockUserRepo(ctrl)

		//when
		res := NewUserService(mUr)
		expected := &userService{mUr}

		//then
		if diff := cmp.Diff(res, expected, cmpopts.IgnoreUnexported(userService{})); diff != "" {
			t.Errorf("Diff: (-got +want)\n%s", diff)
		}
	})
}
