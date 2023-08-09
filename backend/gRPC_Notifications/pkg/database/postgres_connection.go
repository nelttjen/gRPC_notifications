package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"notification_grpc/pkg/env"
)

var _ PostgresConnection = &connection{}

type NoConnectionError struct{}

type NoTransactionError struct{}

type TransactionExists struct{}

func (e *NoConnectionError) Error() string {
	return fmt.Sprintf("There's no active connection.")
}

func (e *NoTransactionError) Error() string {
	return fmt.Sprintf("There's no active transaction.")
}

func (e *TransactionExists) Error() string {
	return fmt.Sprintf("There's already active transaction.")
}

type connection struct {
	host     string
	port     int
	user     string
	password string
	database string

	DBConnection  *gorm.DB
	DBTransaction *gorm.DB
}

type PostgresConnection interface {
	MakeConnection() error
	NewTransaction() error
	ExecTransaction(operation string, values []interface{}) error
	CommitTransaction(endTransaction bool) error
	RollbackTransaction(endTransaction bool) error

	GetDBConnection() (*gorm.DB, error)
	GetDBTransaction() (*gorm.DB, error)
}

func NewPostgresConnection() PostgresConnection {
	newEnv := env.NewEnv()
	host, _ := newEnv.GetEnvAsString("POSTGRES_HOST", "127.0.0.1")
	port, _ := newEnv.GetEnvAsInt("POSTGRES_PORT", 5432)
	user, _ := newEnv.GetEnvAsString("POSTGRES_USER", "postgres")
	password, _ := newEnv.GetEnvAsString("POSTGRES_PASSWORD", "postgres")
	database, _ := newEnv.GetEnvAsString("POSTGRES_DB", "default")

	conn := connection{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
	return &conn
}

func (c *connection) formatConnection(redactPassword bool) string {
	password := "<HIDDEN>"

	if !redactPassword {
		password = c.password
	}

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		c.host, c.user, password, c.database, c.port)
}

func (c *connection) MakeConnection() error {
	log.Printf("INFO: Connecting to postgres database using auth string: %s\n", c.formatConnection(true))

	conn, err := gorm.Open(postgres.Open(c.formatConnection(false)), &gorm.Config{})
	if err != nil {
		return err
	}
	c.DBConnection = conn

	log.Printf("INFO: Connected to postgres database\n")

	return nil
}

func (c *connection) NewTransaction() error {
	if c.DBConnection == nil {
		return &NoConnectionError{}
	}

	if c.DBTransaction != nil {
		return &TransactionExists{}
	}

	c.DBTransaction = c.DBConnection.Begin()

	return nil
}

func (c *connection) CommitTransaction(endTransaction bool) error {
	if c.DBTransaction == nil {
		return &NoTransactionError{}
	}

	c.DBTransaction.Commit()

	if endTransaction {
		c.DBTransaction = nil
	}
	return nil
}

func (c *connection) RollbackTransaction(endTransaction bool) error {
	if c.DBTransaction == nil {
		return &NoTransactionError{}
	}

	c.DBTransaction.Rollback()

	if endTransaction {
		c.DBTransaction = nil
	}
	return nil
}

func (c *connection) ExecTransaction(operation string, values []interface{}) error {
	if c.DBTransaction == nil {
		return &NoTransactionError{}
	}
	c.DBTransaction.Exec(operation, values...)
	if c.DBTransaction.Error != nil {
		return c.DBTransaction.Error
	}
	return nil
}

func (c *connection) GetDBConnection() (*gorm.DB, error) {

	if c.DBConnection == nil {
		return nil, &NoConnectionError{}
	}
	return c.DBConnection, nil
}

func (c *connection) GetDBTransaction() (*gorm.DB, error) {
	if c.DBTransaction == nil {
		return nil, &NoTransactionError{}
	}
	return c.DBTransaction, nil
}
