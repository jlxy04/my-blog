package config

const DriverName = "mysql"

//数据库配置
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

//主库配置
var MasterDbConfig = DbConfig{
	Host:     "127.0.0.1",
	Port:     3306,
	User:     "root",
	Password: "123456",
	DbName:   "blog",
}
