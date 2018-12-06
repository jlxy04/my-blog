package models

import "github.com/my-blog/extend"

type BlogComment struct {
	Id          string      `xorm:"not null pk VARCHAR(40)" json:"-"`
	ArticleId   string      `xorm:"not null comment('文章ID') VARCHAR(40)" json:"-"`
	Commentator string      `xorm:"not null comment('评价人') VARCHAR(40)" json:"commentator"`
	Content     string      `xorm:"not null VARCHAR(255)" json:"content"`
	CreateTime  extend.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME <-" json:"createTime"`
}
