package db

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
)

var pool *sql.DB // Database connection pool.

// Init init pool
func Init(connStr string, logger log.Logger) *sql.DB {
	logger.Log("connection string", connStr)
	var err error
	pool, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Log("unable to use connection string", err)
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	return pool
}

// Ping ping db
func Ping(ctx context.Context, logger log.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		logger.Log("status", "unable to connect to database", "error", err)
	} else {
		logger.Log("status", "database connected")
	}
}

func Query(ctx context.Context, logger log.Logger) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	rows, err := pool.QueryContext(ctx, "select nativeid,name from cloud_vm where name = $1", "HADOOP4")

	if err != nil {
		logger.Log("error", err)
		return
	}
	var names []string
	for rows.Next() {
		var nativeid, name string
		err := rows.Scan(&nativeid, &name)

		if err != nil {
			break
		}
		names = append(names, name, nativeid)
	}
	if closeErr := rows.Close(); closeErr != nil {
		logger.Log("error", err)
		return
	}

	// Check for row scan error.
	if err != nil {
		logger.Log("error", err)
		return
	}

	// Check for errors during row iteration.
	if err = rows.Err(); err != nil {
		logger.Log("error", err)
		return
	}
	logger.Log("result", strings.Join(names, ", "))
}
