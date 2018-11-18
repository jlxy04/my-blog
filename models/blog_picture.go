package models

type BlogPicture struct {
	Id  string `xorm:"not null pk VARCHAR(40)" json:"id"`
	Url string `xorm:"not null VARCHAR(255)" json:"url"`
	Hot int    `xorm:"not null default 0 INT(255)" json:"hot"`
}
