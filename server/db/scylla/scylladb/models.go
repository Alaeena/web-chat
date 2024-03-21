package scylladb

import (
	"github.com/gocql/gocql"
	"time"
)

type Report struct {
	Id        gocql.UUID
	UserId    string
	Type      string
	Message   string
	CreatedAt time.Time
}
