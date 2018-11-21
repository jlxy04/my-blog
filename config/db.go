package config

const DriverName = "mysql"

//数据库配置
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	ShowLog  bool
	DebugLog bool
}

//主库配置
var MasterDbConfig = DbConfig{
	Host:     "120.76.157.34",
	Port:     3306,
	User:     "root",
	Password: "dzy_wx_123456",
	DbName:   "blog",
	ShowLog:  true,
	DebugLog: true,
}
