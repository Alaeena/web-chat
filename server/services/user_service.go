package services

import (
	"context"
	"fmt"
	"server/db/database"
	"server/models"
	"server/utils"
)

type UserService struct {
	queries *database.Queries
}

func GetUser(queries *database.Queries) UserService {
	return UserService{queries: queries}
}

func (s *UserService) RegisterUser(ctx context.Context, payload models.RegisterReq) (*models.UserRes, error) {
	hashed, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("couldn't hash password: %s", err)
	}

	user, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		Username: payload.Username,
		Password: hashed,
		Email:    payload.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %s", err)
	}
	accessToken, err := utils.CreateJwtToken(user)
	if err != nil {
		return nil, fmt.Errorf("error creating access token: %s", err)
	}
	return models.MapUser(user, accessToken), nil
}
func (s *UserService) LoginUser(ctx context.Context, payload models.LoginReq) (*models.UserRes, error) {
	user, err := s.queries.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, fmt.Errorf("no user email found: %s", err)
	}
	err = utils.CheckPassword(payload.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid password: %s", err)
	}
	accessToken, err := utils.CreateJwtToken(user)
	if err != nil {
		return nil, fmt.Errorf("Error creating acces token: %s", err)
	}
	return models.MapUser(user, accessToken), nil
}
