package svc

import (
	"mall/service/user/model"
	"mall/service/user/rpc/internal/config"
	"mall/service/user/rpc/userclient"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel

	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
