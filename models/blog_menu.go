package models

import (
	"time"
)

type BlogMenu struct {
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
	Id         string    `xorm:"not null pk comment('主键') VARCHAR(40)"`
	Name       string    `xorm:"not null comment('菜单名称') VARCHAR(20)"`
	Sort       int       `xorm:"not null comment('菜单排序') INT(11)"`
	Status     string    `xorm:"not null default 'Y' comment('是否启动，Y-启用，N-停用') VARCHAR(1)"`
	Url        string    `xorm:"not null comment('菜单url') VARCHAR(50)"`
}
