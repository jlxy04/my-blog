package models

type BlogCount struct {
	Id   string `xorm:"not null pk VARCHAR(40)" json:"id"`
	Key  string `xorm:"not null comment('统计key') VARCHAR(40)" json:"Key"`
	Val  string `xorm:"not null comment('统计值') VARCHAR(40)" json:"val"`
	Desc string `xorm:"not null comment('统计值') VARCHAR(40)" json:"desc"`
	//AllReadCount  int    `xorm:"not null all_read_count('总阅读量') int" json:"allReadCount"`
	//AllVisitCount int    `xorm:"not null all_visit_count('总访问量') int"  json:"allVisitCount"`
	//AllLikeCount  int    `xorm:"default null all_like_count('总点赞数') int" json:"coverImgUrl"`
}
