// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package postgresdb

import (
	"database/sql"
)

type User struct {
	ID           int32
	Username     string
	Email        string
	Password     string
	AccessToken  sql.NullString
	RefreshToken sql.NullString
}