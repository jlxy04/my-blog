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
	Host:     "",
	Port:     10,
	User:     "",
	Password: "",
	DbName:   "blog",
	ShowLog:  true,
	DebugLog: true,
}
