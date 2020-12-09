package api

import (
	"datacache/models"
	"log"
	"time"
)

var restorequeue chan *models.Cache

func init() {
	restorequeue = make(chan *models.Cache, 20)
}

func (svc *Service) restoreCache() {
	for elem := range restorequeue {
		svc.Memory.Store(*elem.Key, elem)
		svc.Timestamp.Store(*elem.Key, time.Now().UTC())
		log.Println("System: Key Restored:", *elem.Key)
	}
}
