package core

import (
	redisc "github.com/yalvinz/go-helper/redis/redisc"
	myconfig "github.com/yalvinz/go-redis-sample/config"
)

type CoreModule struct {
	RediscCache *redisc.Cluster
}

func NewCoreModule(hitterCfg *myconfig.HitterConfig) (*CoreModule, error) {

	// Initialize redis cluster pool
	rediscCache, err := redisc.InitCluster(&hitterCfg.Redisc)
	if err != nil {
		return nil, err
	}

	return &CoreModule{
		RediscCache: rediscCache,
	}, nil
}
