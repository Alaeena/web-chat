package scylladb

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"os"
)

type Manager struct {
	ScyllaHost     string
	ScyllaKeyspace string
}

func NewManager() *Manager {
	return &Manager{
		ScyllaHost:     os.Getenv("SCYLLA_HOST"),
		ScyllaKeyspace: os.Getenv("SCYLLA_KEYSPACE"),
	}
}
func connect(keyspace string, host string) (gocqlx.Session, error) {
	c := gocql.NewCluster(host)
	c.Keyspace = keyspace
	return gocqlx.WrapSession(c.CreateSession())
}
func (manager *Manager) Connect() (gocqlx.Session, error) {
	return connect(manager.ScyllaKeyspace, manager.ScyllaHost)
}

func (manager *Manager) CreateKeySPace() error {
	session, err := connect("system", manager.ScyllaHost)
	if err != nil {
		return err
	}
	defer session.Close()

	stmt := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`, manager.ScyllaKeyspace)
	return session.ExecStmt(stmt)
}
