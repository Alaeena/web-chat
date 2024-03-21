package scylladb

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/scylladb/gocqlx/v2/qb"
	"log"
	"time"
)

type NewReportsParams struct {
	UserId  string
	Type    string
	Message string
}

func (q *Queries) NewReport(arg *NewReportsParams) {
	tableName := fmt.Sprintf("%s.reports", q.keyspace)
	id, err := gocql.ParseUUID(uuid.New().String())
	if err != nil {
		log.Fatal("Error creating id:", err)
	}
	report := &Report{
		Id:        id,
		UserId:    arg.UserId,
		Type:      arg.Type,
		Message:   arg.Message,
		CreatedAt: time.Now().UTC(),
	}

	stmt := qb.Insert(tableName).Columns("id", "user_id", "type", "message", "created_at").Query(q.session)
	stmt.BindStruct(report)

	if err := stmt.ExecRelease(); err != nil {
		log.Fatal("ExecRelease() failed:", err)
	}
}
