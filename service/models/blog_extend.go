package models

type BlogExtend struct {
	Id              string `xorm:"not null pk VARCHAR(40)"`
	ArticleId       string `xorm:"NOT NULL comment('文章ID') VARCHAR(40)" json:"articleId"`
	ReadCount       int    `xorm:"not null default 0 comment('阅读数量') INT(255)" json:"readCount"`
	LikeCount       int    `xorm:"not null default 0 INT(255)" json:"likeCount"`
	PreArticle      string `xorm:"VARCHAR(40)" json:"preArticle"`
	NexArticle      string `xorm:"VARCHAR(40)" json:"nextArticle"`
	PreArticleName  string `xorm:"pre_article_name " json:"preArticleName"`
	NextArticleName string `xorm:"next_article_name" json:"nextArticleName"`
}
