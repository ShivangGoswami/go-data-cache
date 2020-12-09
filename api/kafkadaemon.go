package api

import "datacache/kafka"

var notifyqueue chan string

func init() {
	notifyqueue = make(chan string, 20)
}

func (svc *Service) notifyKafka() {
	for elem := range notifyqueue {
		kafka.Publish(elem)
	}
}
