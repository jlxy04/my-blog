package models

import "my-blog/extend"

type BlogArticle struct {
	Id           string      `xorm:"not null pk VARCHAR(40)" json:"id"`
	Title        string      `xorm:"not null comment('标题') VARCHAR(40)" json:"title"`
	Author       string      `xorm:"not null comment('作者') VARCHAR(50)"  json:"author"`
	CoverImgUrl  string      `xorm:"default null comment('作者') VARCHAR(50)" json:"coverImgUrl"`
	CategoryId   string      `xorm:"not null comment('分类ID') VARCHAR(40)" json:"categoryId"`
	Label        string      `xorm:"comment('标签') VARCHAR(50)" json:"label"`
	Introduction string      `xorm:"comment('简介') VARCHAR(100)" json:"introduction"`
	Content      string      `xorm:"not null comment('内容') TEXT" json:"content"`
	CreateTime   extend.Time `xorm:"not null default 'CURRENT_TIMESTAMP' DATETIME" json:"createTime"`
}
