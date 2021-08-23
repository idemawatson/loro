package models

type User struct {
	Id                 string `dynamo:"id"`
	UserId             string `dynamo:"userId"`
	UserName           string `dynamo:"userName"`
	Description        string `dynamo:"description"`
	MailAddress        string `dynamo:"mailAddress"`
	ProfilePicturePath string `dynamo:"profilePicturePath"`
	NumPost            int    `dynamo:"numPost"`
	NumFollowing       int    `dynamo:"numFollowing"`
	NumFollowed        int    `dynamo:"numFollowed"`
	NumFavorite        int    `dynamo:"numFavorite"`
}

type GetUserInfoRequest struct {
	UserId string `json:"userId"`
}
