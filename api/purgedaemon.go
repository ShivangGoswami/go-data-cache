package api

import (
	"datacache/models"
	"log"
	"os"
	"time"
)

var sleeptime time.Duration

func init() {
	if temp := os.Getenv("keeperSleepTime"); temp != "" {
		if dur, err := time.ParseDuration(temp); err == nil {
			log.Println("System: Keeper Sleep Time:", dur)
			sleeptime = dur
			return
		} else {
			log.Println("System: Keeper Sleep Time error:", err)
		}
	}
	log.Println("System: Keeper Sleep Time not Defined Using default 1 minute")
	sleeptime = time.Minute
}
func (svc *Service) gatekeeper() {
	for {
		time.Sleep(sleeptime)
		//log.Println("System: Purge Daemon process Started")
		var elapsedKey []interface{}
		svc.Memory.Range(func(key, value interface{}) bool {
			if temptime, ok := svc.Timestamp.Load(key); ok {
				if tstamp, ok := temptime.(time.Time); ok {
					if cache, ok := value.(*models.Cache); ok {
						if elapsedTime, err := time.ParseDuration(*cache.Expiration); err == nil {
							if time.Now().UTC().Sub(tstamp) > elapsedTime {
								elapsedKey = append(elapsedKey, key)
							}
						}
					}
				}
			}
			return true
		})
		for _, val := range elapsedKey {
			log.Println("System: Key Purged:", val)
			svc.Memory.Delete(val)
			svc.Timestamp.Delete(val)
		}
		//log.Println("System: Purge Daemon process Completed")
	}
}
