package models

import (
	"my-blog/extend"
)

type BlogLeaveMsg struct {
	Id           string      `xorm:"not null pk VARCHAR(40)"`
	Name         string      `xorm:"not null comment('姓名') VARCHAR(50)"`
	Mail         string      `xorm:"not null comment('邮箱') VARCHAR(255)"`
	Content      string      `xorm:"not null comment('留言内容') VARCHAR(255)"`
	CreateTime   extend.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	ReplyContent string      `xorm:"comment('回复内容') VARCHAR(255)"`
	ReplyTime    extend.Time `xorm:"comment('回复时间') DATETIME"`
}
