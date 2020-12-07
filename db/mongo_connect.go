package db

import (
	"context"
	"datacache/restapi/ops/data_cache_operations"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatastore struct {
	client *mongo.Client
}

type Datastore interface {
	Disconnect()
	StoreCacheObj(params data_cache_operations.PostStoreParams) error
}

var conn struct {
	ConnectionURI   string
	MongoDatabase   string
	MongoCollection string
}

func init() {
	if temp := os.Getenv("mongohost"); temp != "" {
		conn.ConnectionURI = temp
	} else {
		log.Fatal("System:", "mongo host not defined")
	}
	if temp := os.Getenv("mongoDatabase"); temp != "" {
		conn.MongoDatabase = temp
	} else {
		log.Fatal("System:", "mongo database not defined")
	}
	if temp := os.Getenv("mongoCollection"); temp != "" {
		conn.MongoCollection = temp
	} else {
		log.Fatal("System:", "mongo collection not defined")
	}
	log.Printf("System/MongoMetadata:%v", conn)
}

func (mds *MongoDatastore) Disconnect() {
	log.Println("Mongo/System: Disconnect:", mds.client.Disconnect(context.TODO()))
}

func Connect() Datastore {
	opts := options.Client().ApplyURI(conn.ConnectionURI)
	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Fatal("Mongo/System:", err)
	}
	if err = client.Connect(context.TODO()); err != nil {
		log.Fatal("Mongo/System:", err)
	}
	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Mongo/System:", err)
	}
	log.Println("Mongo/System:", "Connected")
	return &MongoDatastore{client: client}
}
