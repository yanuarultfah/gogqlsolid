package db

import (
	"database/sql"
	"fmt"
	"gogqlgensolid/internal/config"

	"github.com/XSAM/otelsql"
	_ "github.com/lib/pq" // PostgreSQL driver
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

var DB *sql.DB

func Init(cfg config.PostgresConfig) {
	driverName, errotel := otelsql.Register("postgres",
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithSQLCommenter(true),
	)
	if errotel != nil {
		panic("Failed to register otelsql: " + errotel.Error())
	}
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	DB, err = sql.Open(driverName, connStr)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	DB.SetMaxOpenConns(cfg.MaxOpenConns)
	DB.SetMaxIdleConns(cfg.MaxIdleConns)

	if err = DB.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	}
}
