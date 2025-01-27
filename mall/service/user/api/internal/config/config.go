package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserRpc zrpc.RpcClientConf

	Mysql struct {
		DataSource string
	}

	CacheRedis []struct {
		Host string
		Pass string
		Type string
	}

	Salt string
}
