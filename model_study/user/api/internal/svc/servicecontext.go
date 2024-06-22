package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zeroStudy/model_study/user/api/internal/config"
	"zeroStudy/model_study/user/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(mysqlConn),
	}
}
