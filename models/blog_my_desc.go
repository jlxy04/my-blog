package models

type BlogMyDesc struct {
	Id     string `xorm:"not null pk VARCHAR(40)"`
	Name   string `xorm:"not null comment('姓名') VARCHAR(10)"`
	MyDesc string `xorm:"not null comment('自我描述') VARCHAR(255)"`
}
