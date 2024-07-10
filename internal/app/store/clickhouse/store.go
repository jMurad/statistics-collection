package clickhouse

import (
	clickhouse "github.com/ClickHouse/clickhouse-go/v2//"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

// Store ...
type Store struct {
	conn driver.Conn
	// userRepository *UserRepository
}

// New ...
func New(conn driver.Conn) *Store {
	clickhouse.NewClickhouseWriter
	return &Store{
		conn: conn,
	}
}
