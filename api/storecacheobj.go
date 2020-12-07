package api

import (
	"datacache/customerror"
	"datacache/models"
	"datacache/restapi/ops/data_cache_operations"
	"log"
	"time"
)

func (svc *Service) StoreData(params data_cache_operations.PostStoreParams) error {
	log.Println("Request Received in API Layer/key:", *params.Params.Key)
	err := checkDataValues(params.Params)
	if err != nil {
		log.Println("Parsing Error:", err)
		return customerror.TimeParseError
	}
	svc.Timestamp.Store(*params.Params.Key, time.Now().UTC())
	svc.Memory.Store(*params.Params.Key, params.Params)
	err = svc.Store.StoreCacheObj(params)
	if err != nil {
		log.Println("Mongo Error:", err)
		return customerror.MongoStoreException
	}
	log.Println("Request Completed in API Layer/key:", *params.Params.Key)
	return nil
}

func checkDataValues(params *models.Cache) error {
	_, err := time.ParseDuration(*params.Expiration)
	if err != nil {
		return err
	}
	return nil
}
