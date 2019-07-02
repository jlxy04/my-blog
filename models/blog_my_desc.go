package models

type BlogMyDesc struct {
	Id       string `xorm:"not null pk VARCHAR(40)"`
	Name     string `xorm:"not null comment('姓名') VARCHAR(10)" json:"name"`
	MyDesc   string `xorm:"not null comment('自我描述') VARCHAR(255)" json:"myDesc"`
	BlogDesc string `xorm:"not null comment('网站描述') VARCHAR(100)" json:"blogDesc"`
	Icp      string `xorm:"not null comment('网站备案号') VARCHAR(50)" json:"icp"`
}
