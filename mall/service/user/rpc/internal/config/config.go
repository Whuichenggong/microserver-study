package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RestConf rest.RestConf      `yaml:"RestConf"`
	RpcConf  zrpc.RpcServerConf `yaml:"RpcServerConf"`

	Mysql struct {
		DataSource string `yaml:"DataSource"`
	}

	CacheRedis cache.CacheConf `yaml:"CacheRedis"`
	Salt       string          `yaml:"Salt"`

	Auth struct {
		AccessSecret string `yaml:"AccessSecret"`
		AccessExpire int64  `yaml:"AccessExpire"`
	}

	UserRpc zrpc.RpcClientConf `yaml:"UserRpc"`

	Mode string `yaml:"Mode,default=dev"` // 指定默认值

	ListenOn string `yaml:"ListenOn"` // 指定默认值
}
