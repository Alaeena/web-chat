package user

import (
	"context"
	"database/sql"
	"fmt"
	"server/db/postgres/postgresdb"
	"server/utils"
)

type postgresService struct {
	queries *postgresdb.Queries
}

func GetPostgresService(queries *postgresdb.Queries) Service[postgresdb.User] {
	return &postgresService{queries: queries}
}

func (s *postgresService) createToken(ctx context.Context, user postgresdb.User) (string, error) {
	accessToken, err := utils.CreateJwtToken(string(user.ID), user.Email)
	if err != nil {
		return "", err
	}
	err = s.queries.SetAccessToken(ctx, postgresdb.SetAccessTokenParams{
		AccessToken: sql.NullString{String: accessToken, Valid: true},
		ID:          user.ID,
	})
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *postgresService) RegisterUser(ctx context.Context, payload RegisterReq) (*Response, error) {
	hashed, err := utils.HashPassword(payload.Password)
	if err != nil {
		return nil, fmt.Errorf("couldn't hash password: %s", err)
	}

	user, err := s.queries.CreateUser(ctx, postgresdb.CreateUserParams{
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
	return MapPostgresUser(user, accessToken), nil
}
func (s *postgresService) LoginUser(ctx context.Context, payload LoginReq) (*Response, error) {
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
	return MapPostgresUser(user, accessToken), nil
}
func (s *postgresService) Logout(ctx context.Context, user postgresdb.User) error {
	err := s.queries.SetAccessToken(ctx, postgresdb.SetAccessTokenParams{
		ID:          user.ID,
		AccessToken: sql.NullString{Valid: false},
	})
	return err
}
