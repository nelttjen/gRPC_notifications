package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	cfg "notification_grpc/internal/config"
	myenv "notification_grpc/pkg/env"
)

var _ MongoClient = &mongoClient{}

type NoMongoConnectionError struct{}

func (e *NoMongoConnectionError) Error() string {
	return "No active mongo connection to close"
}

type MongoClient interface {
	MakeConnection() error
	CloseConnection() error
	GetCollectionFromAdmin(collectionName string) (*mongo.Collection, error)
}

type mongoClient struct {
	Connection *mongo.Client
}

func NewMongoConnection() MongoClient {
	return &mongoClient{}
}

func (c *mongoClient) MakeConnection() error {
	env := myenv.NewEnv()
	err := env.LoadEnv(cfg.EnvRoot)
	if err != nil {
		return err
	}

	host, _ := env.GetEnvAsString("MONGODB_HOST", "127.0.0.1")
	port, _ := env.GetEnvAsInt("MONGODB_PORT", 27017)
	username, _ := env.GetEnvAsString("MONGODB_USERNAME", "admin")
	password, _ := env.GetEnvAsString("MONGODB_PASSWORD", "adminpass123")
	database, _ := env.GetEnvAsString("MONGODB_AUTHSOURCE", "development")

	authString := fmt.Sprintf("mongodb://%s:%d", host, port)

	log.Println("Connecting to MongoDB using auth string: ", authString)

	clientOptions := options.Client().ApplyURI(authString).SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    database,
		Username:      username,
		Password:      password,
		PasswordSet:   false,
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	c.Connection = client
	return nil
}

func (c *mongoClient) CloseConnection() error {
	if c.Connection == nil {
		return &NoMongoConnectionError{}
	}

	if err := c.Connection.Disconnect(context.Background()); err != nil {
		return err
	}
	return nil
}

func (c *mongoClient) GetCollectionFromAdmin(collectionName string) (*mongo.Collection, error) {
	if c.Connection == nil {
		return nil, &NoMongoConnectionError{}
	}

	return c.Connection.Database("admin").Collection(collectionName), nil
}
