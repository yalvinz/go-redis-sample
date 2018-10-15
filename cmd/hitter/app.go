package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/yalvinz/go-helper/grace"
	"github.com/yalvinz/go-helper/logging"
	myconfig "github.com/yalvinz/go-redis-sample/config"
	"github.com/yalvinz/go-redis-sample/core"
)

func main() {
	var err error

	// Initialize logger
	flag.Parse()
	logging.LogInit()

	// Initialize config
	serviceName := "hitter"
	hitterCfg := &myconfig.HitterConfig{}
	if err = myconfig.InitConfig(hitterCfg, serviceName, "files/etc/go-redis-sample"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("loaded %s service config is: %+v\n", serviceName, hitterCfg)

	// Initialize core module
	coreModule, err := core.NewCoreModule(hitterCfg)
	if err != nil {
		log.Fatal("Error initializing core module, err:", err)
	}

	// Initialize handlers
	r := mux.NewRouter()
	r.HandleFunc("/redisc/status", coreModule.GetClusterStatus)
	r.HandleFunc("/redisc/del/{key}", coreModule.DoRedisDelKey)
	r.HandleFunc("/redisc/get/{key}", coreModule.DoRedisGetKey)
	r.HandleFunc("/redisc/setex/{key}/{value}/{ttl}", coreModule.DoRedisSetexKey)
	r.HandleFunc("/redisc/hget/{key}/{field}", coreModule.DoRedisHGetKey)
	r.HandleFunc("/redisc/hset/{key}/{field}/{value}/{ttl}", coreModule.DoRedisHSetKey)
	r.HandleFunc("/redisc/hmget/{key}", coreModule.DoRedisHMGetKey)
	r.HandleFunc("/redisc/hmset/{key}/{ttl}", coreModule.DoRedisHMSetKey)

	log.Fatal(grace.Serve(":9000", r))
}
