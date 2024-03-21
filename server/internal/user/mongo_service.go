package user

import (
	"context"
	"fmt"
	"server/db/mongo/mongodb"
	"server/utils"
)

type mongoService struct {
	queries *mongodb.Queries
}

func GetMongoService(queries *mongodb.Queries) Service[mongodb.User] {
	return &mongoService{queries: queries}
}

func (s *mongoService) createToken(ctx context.Context, user mongodb.User) (string, error) {
	accessToken, err := utils.CreateJwtToken(user.ID.Hex(), user.Email)
	if err != nil {
		return "", err
	}
	err = s.queries.SetAccessToken(ctx, mongodb.SetAccessTokenParams{
		AccessToken: accessToken,
		ID:          user.ID.Hex(),
	})
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *mongoService) RegisterUser(ctx context.Context, payload RegisterReq) (*Response, error) {
	hashed, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("couldn't hash password: %s", err)
	}

	user, err := s.queries.CreateUser(ctx, mongodb.CreateUserParams{
		Username: payload.Username,
		Password: hashed,
		Email:    payload.Email,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating user: %s", err)
	}
	accessToken, err := s.createToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error creating access token: %s", err)
	}
	return MapMongoUser(user, accessToken), nil
}
func (s *mongoService) LoginUser(ctx context.Context, payload LoginReq) (*Response, error) {
	user, err := s.queries.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return nil, fmt.Errorf("no user email found: %s", err)
	}
	err = utils.CheckPassword(payload.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid password: %s", err)
	}
	accessToken, err := s.createToken(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error creating access token: %s", err)
	}
	return MapMongoUser(user, accessToken), nil
}
func (s *mongoService) Logout(ctx context.Context, user mongodb.User) error {
	err := s.queries.SetAccessToken(ctx, mongodb.SetAccessTokenParams{
		ID:          user.ID.Hex(),
		AccessToken: "",
	})
	return err
}
