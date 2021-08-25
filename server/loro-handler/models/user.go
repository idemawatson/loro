package models

type User struct {
	Id                 string `dynamo:"id"`
	UserId             string `dynamo:"userId" json:"userId"`
	UserName           string `dynamo:"userName" json:"userName"`
	Description        string `dynamo:"description" json:"description"`
	MailAddress        string `dynamo:"mailAddress" json:"mailAddress"`
	ProfilePicturePath string `dynamo:"profilePicturePath" json:"profilePicturePath"`
	NumPost            int    `dynamo:"numPost" json:"numPost"`
	NumFollowing       int    `dynamo:"numFollowing" json:"numFollowing"`
	NumFollowed        int    `dynamo:"numFollowed" json:"numFollowed"`
	NumFavorite        int    `dynamo:"numFavorite" json:"numFavorite"`
}

type GetUserInfoRequest struct {
	UserId string `json:"userId"`
}
