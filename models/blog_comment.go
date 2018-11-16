package models

import (
	"time"
)

type BlogComment struct {
	Id          string    `xorm:"not null pk VARCHAR(40)"`
	ArticleId   string    `xorm:"not null comment('文章ID') VARCHAR(40)"`
	Commentator string    `xorm:"not null comment('评价人') VARCHAR(40)"`
	Content     string    `xorm:"not null VARCHAR(255)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME"`
}
