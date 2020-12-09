package db

import (
	"datacache/customerror"
	"datacache/models"
	"datacache/restapi/ops/data_cache_operations"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mds *MongoDatastore) Fetchkey(params data_cache_operations.GetFetchParams) (*models.Cache, error) {
	log.Println("Fetch Request Received in DB Layer/key:", *params.Key)
	var result models.Cache
	err := mds.client.Database(conn.MongoDatabase).Collection(conn.MongoCollection).FindOne(params.HTTPRequest.Context(), bson.D{{"key", *params.Key}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, customerror.MongoNotFound
		}
		return nil, err
	}
	log.Println("Fetch Request Completed in DB Layer/key:", *params.Key)
	return &result, nil
}
