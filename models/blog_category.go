package models

type BlogCategory struct {
	Id       string `xorm:"not null pk comment('文章分类名') VARCHAR(40)"`
	Name     string `xorm:"not null comment('文章分类名') VARCHAR(40)"`
	Sort     int    `xorm:"not null default 0 comment('排序') INT(255)"`
	ShowMain string `xorm:"not null default 'Y' VARCHAR(1)"`
}
