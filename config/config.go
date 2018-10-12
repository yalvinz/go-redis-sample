package config

import (
	config "github.com/yalvinz/go-helper/config"
	redisc "github.com/yalvinz/go-helper/redis/redisc"
)

type HitterConfig struct {
	Server ServerCfg
	Redisc redisc.Config
}

func InitConfig(cfg interface{}, module string, path ...string) error {
	if len(path) == 0 {
		path = GetDefaultConfigPaths()
	}

	err := config.ReadModuleConfig(cfg, module, path)
	if err != nil {
		return err
	}

	return nil
}

func GetDefaultConfigPaths() []string {
	return []string{"files/etc/go-redis-sample", "/etc/go-redis-sample"}
}
