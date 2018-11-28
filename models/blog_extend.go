package models

type BlogExtend struct {
	Id              string `xorm:"not null pk VARCHAR(40)"`
	ArticleId       string `xorm:"NOT NULL comment('文章ID') VARCHAR(40)"`
	ReadCount       int    `xorm:"not null default 0 comment('阅读数量') INT(255)"`
	LikeCount       int    `xorm:"not null default 0 INT(255)"`
	PreArticle      string `xorm:"VARCHAR(40)"`
	NexArticle      string `xorm:"VARCHAR(40)"`
	PreArticleName  string `json:"preArticleName"`
	NextArticleName string `json:"nextArticleName"`
}
