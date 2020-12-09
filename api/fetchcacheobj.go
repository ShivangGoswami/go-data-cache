package api

import (
	"datacache/customerror"
	"datacache/models"
	"datacache/restapi/ops/data_cache_operations"
	"log"
)

func (svc *Service) FetchData(params data_cache_operations.GetFetchParams) (models.CacheCollection, error) {
	log.Println("Fetch Request Received in API Layer")
	defer log.Println("Fetch Request Completed in API Layer")
	if params.Key != nil {
		value, ok := svc.Memory.Load(*params.Key)
		if ok {
			if temp := value.(*models.Cache); ok {
				return models.CacheCollection{temp}, nil
			}
		}
		dbresult, err := svc.Store.Fetchkey(params)
		if err != nil {
			log.Println("Mongo Error:", err)
			return nil, err
		}
		//notify kafka
		notifyqueue <- *params.Key
		//restore cache
		restorequeue <- dbresult
		return models.CacheCollection{dbresult}, nil
	} else if params.Index != nil {
		//switch to db mode
		response, err := svc.Store.FetchIndex(params)
		if err != nil {
			log.Println("Mongo Error:", err)
			return nil, err
		}
		for _, value := range response {
			if _, ok := svc.Memory.Load(*value.Key); !ok {
				//notify kafka
				notifyqueue <- *value.Key
				//restore cache
				restorequeue <- value
			}
		}
		return response, nil
	} else {
		return nil, customerror.InvalidInput
	}
	//return nil, nil
}

func (svc *Service) getLength() int {
	length := 0
	svc.Memory.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}
