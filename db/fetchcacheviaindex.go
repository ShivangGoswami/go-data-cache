package db

import (
	"datacache/models"
	"datacache/restapi/ops/data_cache_operations"
	"log"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var pagesize int64

func init() {
	if temp := os.Getenv("pagesize"); temp != "" {
		if size, err := strconv.Atoi(temp); err == nil {
			log.Println("System: Page Size:", size)
			pagesize = int64(size)
			return
		} else {
			log.Println("System: Page Size Parser error:", err.Error())
		}
	}
	log.Println("System: using default page size of 10")
	pagesize = 10
}

func (mds *MongoDatastore) FetchIndex(params data_cache_operations.GetFetchParams) (models.CacheCollection, error) {
	log.Println("Fetch Request Received in DB Layer/index:", *params.Index)
	var result models.CacheCollection
	cursor, err := mds.client.Database(conn.MongoDatabase).Collection(conn.MongoCollection).Find(params.HTTPRequest.Context(), bson.D{{}}, options.Find().SetLimit(pagesize).SetSkip((*params.Index-1)*pagesize))
	if err != nil {
		return nil, err
	}
	err = cursor.All(params.HTTPRequest.Context(), &result)
	log.Println("Fetch Request Completed in DB Layer/index:", *params.Index)
	return result, nil
}
