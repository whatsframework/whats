package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"whats/core/config"
)

var (
	// MongoDB MongoDB
	MongoDB *mongo.Database
)

// MongoConn MongoConn
//type MongoConn struct {
//	clientOptions *options.ClientOptions
//	client        *mongo.Client
//	collections   *mongo.Collection
//}

// InitMongoDB  InitMongoDB
func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetDefaultEnv("MONGO_URI", "mongodb://root:123456@127.0.0.1:27017")))
	if err != nil {
		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	MongoDB = client.Database("trace")

}
