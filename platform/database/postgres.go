package database

import (
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
)

// postgresql database와 커넥션 맺는 함수
func PostgreSQLConnection() (*sqlx.DB, error) {
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns =2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) //  0

	// connection은 됬는데 ping이 db로 안가는 에러
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
