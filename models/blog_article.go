package models

import (
	"time"
)

type BlogArticle struct {
	Id           string    `xorm:"not null pk VARCHAR(40)"`
	Title        string    `xorm:"not null comment('标题') VARCHAR(40)"`
	Author       string    `xorm:"not null comment('作者') VARCHAR(50)"`
	CategoryId   string    `xorm:"not null comment('分类ID') VARCHAR(40)"`
	Label        string    `xorm:"comment('标签') VARCHAR(10)"`
	Introduction string    `xorm:"comment('简介') VARCHAR(100)"`
	Content      string    `xorm:"not null comment('内容') TEXT"`
	CreateTime   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
