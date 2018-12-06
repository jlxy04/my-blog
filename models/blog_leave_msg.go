package models

import (
	"github.com/my-blog/extend"
)

type BlogLeaveMsg struct {
	Id           string      `xorm:"not null pk VARCHAR(40)"`
	Name         string      `xorm:"not null comment('姓名') VARCHAR(50)" json:"name" form:"name"`
	Mail         string      `xorm:"not null comment('邮箱') VARCHAR(255)" json:"mail" form:"email"`
	Content      string      `xorm:"not null comment('留言内容') VARCHAR(255)" json:"content" form:"content"`
	CreateTime   extend.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME" json:"createTime"`
	ReplyContent string      `xorm:"comment('回复内容') VARCHAR(255)" json:"replyContent"`
	ReplyTime    extend.Time `xorm:"comment('回复时间') DATETIME" json:"replyTime"`
}
