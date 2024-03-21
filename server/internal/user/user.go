package user

import (
	"server/db/mongo/mongodb"
	"server/db/postgres/postgresdb"
)

type User interface {
	mongodb.User | postgresdb.User
}
type RegisterReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Response struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func MapPostgresUser(model postgresdb.User, token string) *Response {
	return &Response{
		ID:          string(model.ID),
		Username:    model.Username,
		AccessToken: token,
	}
}
func MapMongoUser(model mongodb.User, token string) *Response {
	return &Response{
		ID:          model.ID.Hex(),
		Username:    model.Username,
		AccessToken: token,
	}
}
