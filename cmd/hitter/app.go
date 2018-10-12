package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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
	http.HandleFunc("/redisc/status", coreModule.GetClusterStatus)
	http.HandleFunc("/redisc/get/", coreModule.GetKey)

	log.Fatal(grace.Serve(":9000", nil))
}
