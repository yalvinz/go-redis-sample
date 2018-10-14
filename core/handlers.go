package core

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func (cm *CoreModule) GetClusterStatus(w http.ResponseWriter, r *http.Request) {
	val, _ := cm.RediscCache.ClusterStatus()

	log.Printf("This is stats %s", string(val))

	return
}

func (cm *CoreModule) DoRedisGetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	val, err := cm.RediscCache.Get(key)
	if err != nil {
		log.Printf("Error get key %s, err %s", key, err)
		return
	}

	log.Printf("This is %s val: %s", key, val)

	return
}

func (cm *CoreModule) DoRedisSetexKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	ttl, _ := strconv.Atoi(vars["ttl"])

	err := cm.RediscCache.Setex(key, value, ttl)
	if err != nil {
		log.Printf("Error setex key %s with value %s, err %s", key, value, err)
		return
	}

	log.Printf("Success setex key %s with value %s", key, value)

	return
}

func (cm *CoreModule) DoRedisHGetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	field := vars["field"]

	val, err := cm.RediscCache.HGet(key, field)
	if err != nil {
		log.Printf("Error hget key %s, err %s", key, err)
		return
	}

	log.Printf("This is %s val: %s", key, val)

	return
}

func (cm *CoreModule) DoRedisHSetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	value := vars["value"]
	field := vars["field"]
	ttl, _ := strconv.Atoi(vars["ttl"])

	err := cm.RediscCache.HSet(key, field, value, ttl)
	if err != nil {
		log.Printf("Error hset key %s with value %s, err %s", key, value, err)
		return
	}

	log.Printf("Success hset key %s with value %s", key, value)

	return
}

func (cm *CoreModule) DoRedisHMGetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fields := []string{}
	fields = append(fields, "name")
	fields = append(fields, "age")

	val, err := cm.RediscCache.HMGet(key, fields)
	if err != nil {
		log.Printf("Error hmget key %s, err %s", key, err)
		return
	}

	log.Printf("This is %s val: %s", key, strings.Join(val, ", "))

	return
}

func (cm *CoreModule) DoRedisHMSetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]
	ttl, _ := strconv.Atoi(vars["ttl"])

	m := make(map[string]string)
	m["name"] = "yoshua"
	m["age"] = "18"
	m["gender"] = "male"

	err := cm.RediscCache.HMSet(key, ttl, m)
	if err != nil {
		log.Printf("Error hmset key %s, err %s", key, err)
		return
	}

	log.Printf("Success hmset key %s with value %s", key, m)

	return
}
