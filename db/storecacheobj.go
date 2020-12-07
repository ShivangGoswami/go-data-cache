package db

import (
	"datacache/restapi/ops/data_cache_operations"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (mds *MongoDatastore) StoreCacheObj(params data_cache_operations.PostStoreParams) error {
	log.Println("Request Received in DB Layer/key:", *params.Params.Key)
	result, err := mds.client.Database(conn.MongoDatabase).Collection(conn.MongoCollection).UpdateOne(params.HTTPRequest.Context(), bson.D{{"key", *params.Params.Key}}, bson.D{{"$set", params.Params}}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	log.Printf("Mongo Result:%#v\n", result)
	log.Println("Request Completed in DB Layer/key:", *params.Params.Key)
	return nil
}
