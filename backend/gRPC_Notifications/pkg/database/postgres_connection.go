package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"notification_grpc/pkg/env"
)

var _ Connection = &connection{}

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

type Connection interface {
	MakeConnection() error
	NewTransaction() error
	ExecTransaction(operation string, values []interface{}) error
	CommitTransaction()
	RollbackTransaction(endTransaction bool)

	GetDBConnection() (*gorm.DB, error)
}

func NewConnection() Connection {
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

func (c *connection) formatConnection() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow",
		c.host, c.user, c.password, c.database, c.port)
}

func (c *connection) MakeConnection() error {
	conn, err := gorm.Open(postgres.Open(c.formatConnection()), &gorm.Config{})
	if err != nil {
		return err
	}
	c.DBConnection = conn
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

func (c *connection) CommitTransaction() {
	c.DBTransaction.Commit()

	c.DBTransaction = nil
}

func (c *connection) RollbackTransaction(endTransaction bool) {
	c.DBTransaction.Rollback()

	if endTransaction {
		c.DBTransaction = nil
	}
}

func (c *connection) ExecTransaction(operation string, values []interface{}) error {
	if c.DBTransaction == nil {
		return &NoTransactionError{}
	}
	c.DBTransaction.Exec(operation, values)
	return nil
}

func (c *connection) GetDBConnection() (*gorm.DB, error) {

	if c.DBConnection == nil {
		return nil, &NoConnectionError{}
	}
	return c.DBConnection, nil
}