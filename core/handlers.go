package core

import (
	"log"
	"net/http"
	"strings"
)

func (cm *CoreModule) GetClusterStatus(w http.ResponseWriter, r *http.Request) {
	val, _ := cm.RediscCache.ClusterStatus()

	log.Printf("This is stats %s", string(val))

	return
}

func (cm *CoreModule) GetKey(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/redisc/get/")

	val, err := cm.RediscCache.Get(key)
	if err != nil {
		log.Println("Error get key:", err)
		return
	}

	// result, _ := strconv.Atoi(string(val))
	result := string(val)

	log.Printf("This is %s val: %s", key, result)

	return
}
