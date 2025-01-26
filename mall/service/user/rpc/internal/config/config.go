package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	RpcServerConf zrpc.RpcServerConf `yaml:",inline"` // 使用 inline 标签来避免与配置文件中的字段冲突

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	Salt string

	Mode string `yaml:"Mode,default=dev"` // 指定默认值

	ListenOn string `yaml:"ListenOn"` // 指定默认值
}
