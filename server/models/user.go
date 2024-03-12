package models

import "server/db/database"

type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserRes struct {
	ID          int32  `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func MapUser(model database.User, token string) *UserRes {
	return &UserRes{
		ID:          model.ID,
		Username:    model.Username,
		AccessToken: token,
	}
}
