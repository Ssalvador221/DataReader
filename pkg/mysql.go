package pkg

import (
	"database/sql"
	"fmt"
	"go-data-read/internal/db/store"
	"log"


	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	Conn    *sql.DB
	Queries *store.Queries
}

func NewDB() (*DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/csvreader"
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Ops! Got an error: %w", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("Ops! Got an error: %w", err)
	}

	queries := store.New(conn)

	return &DB{
		Conn:    conn,
		Queries: queries,
	}, nil
}

func (db *DB) Close() {
	if err := db.Conn.Close(); err != nil {
		log.Fatalf("failed to close DB connection: %v", err)
	}
}

func NewNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: v, Valid: true}
}

func NewNullInt32(v int32) sql.NullInt32 {
	if v == 0 {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: v, Valid: true}
}
