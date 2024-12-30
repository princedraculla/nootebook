package utils

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func PostgresConn(host, port, user, password, dbname, sslmode string, maxopenconn, maxidleconn int, timeout time.Duration) (*sql.DB, error) {
	connString := PostgresURI(host, port, user, password, dbname, sslmode)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(maxopenconn)
	conn.SetMaxIdleConns(maxidleconn)
	dbContext, _ := context.WithTimeout(context.Background(), timeout)
	err = conn.PingContext(dbContext)
	if err != nil {
		return nil, fmt.Errorf("error in pinging postgres database: %w", err)
	}
	return conn, nil
}
func PostgresURI(host, port, user, pass, database, sslmode string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, pass, database, sslmode)
}
