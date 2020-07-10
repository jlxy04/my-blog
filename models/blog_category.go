package models

type BlogCategory struct {
	Id       string `xorm:"not null pk comment('主键') VARCHAR(40)" json:"id"`
	Name     string `xorm:"not null comment('文章分类名') VARCHAR(40)" json:"name"`
	Sort     int    `xorm:"not null default 0 comment('排序') INT(255)" json:"sort"`
	ShowMain string `xorm:"not null default 'Y' VARCHAR(1)"`
	Count    int    `xorm:"not null default 0 comment('文章数量') INT(255)" json:"count"`
}
