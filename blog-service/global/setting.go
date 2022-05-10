package global

import (
	"GoprogrammingTourBook/blog-service/pkg/logger"
	"GoprogrammingTourBook/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettindS
)
