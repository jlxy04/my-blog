package datasource

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"my-blog/config"
	"sync"
)

//root:123456@/blog?charset=utf8
var (
	mastEngine  *xorm.Engine
	slaveEngine *xorm.Engine
	lock        sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	if mastEngine != nil {
		return mastEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if mastEngine != nil {
		return mastEngine
	}
	c := config.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Password,
		c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(config.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.InstanceMaster err = ", err)
		return nil
	} else {
		mastEngine = engine
		return mastEngine
	}
}

func InstanceSlave() *xorm.Engine {
	// TODO:
	return slaveEngine
}
