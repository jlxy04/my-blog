package models

type BlogPicture struct {
	Id  string `xorm:"not null pk VARCHAR(40)"`
	Url string `xorm:"not null VARCHAR(255)"`
	Hot int    `xorm:"not null default 0 INT(255)"`
}
