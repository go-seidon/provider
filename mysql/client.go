package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"context"

	_ "github.com/go-sql-driver/mysql"
)

type Client interface {
	Pingable
	Manageable
	Prepareable
	Queryable
	Executable
	Beginable
}

type Pingable interface {
	PingContext(ctx context.Context) error
	Ping() error
}

type Manageable interface {
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	SetConnMaxIdleTime(d time.Duration)
}

type Beginable interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

type Transaction interface {
	Commit() error
	Rollback() error
	Prepareable
	Queryable
	Executable
}

type Prepareable interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
}

type Queryable interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Executable interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

func NewClient(opts ...ClientOption) (*sql.DB, error) {
	p := ClientParam{}
	for _, opt := range opts {
		opt(&p)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		p.Username, p.Password,
		p.Host, p.Port, p.DbName,
	)

	if p.ShouldParseTime {
		dsn = fmt.Sprintf("%s?parseTime=true", dsn)
	}
	return sql.Open("mysql", dsn)
}
