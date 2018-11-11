package datasource

import (
	"github.com/go-xorm/xorm"
	"sync"
)

//root:123456@/blog?charset=utf8
var (
	mastEngine  *xorm.Engine
	slaveEngine *xorm.Engine
	lock        sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	// TODO:
	return mastEngine
}

func InstanceSlave() *xorm.Engine {
	// TODO:
	return slaveEngine
}
