package api

import (
	"datacache/db"
	"log"
	"sync"
)

type Service struct {
	Store     db.Datastore
	Memory    sync.Map
	Timestamp sync.Map
}

func NewService() *Service {
	log.Println("System Started Creating Service Object")
	var svc Service
	svc.Store = db.Connect()
	//purge daemon
	go svc.gatekeeper()
	//restore daemon
	go svc.restoreCache()
	//kafka daemon
	go svc.notifyKafka()
	log.Println("System Completed Creating Service Object")
	return &svc
}

func (svc *Service) DestroyService() {
	log.Println("System Started Destroying Service Object")
	svc.Store.Disconnect()
	log.Println("System Completed Destroying Service Object")
}
