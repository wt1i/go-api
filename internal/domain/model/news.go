package model

type NewsStatus string

const (
	NewsStatusDelete  = "deleted"
	NewsStatusDraft   = "draft"
	NewsStatusPublish = "publish"
)

// News represent entity of the news
type News struct {
	BaseModel
	Title   string     `json:"title"`
	Slug    string     `json:"slug"`
	Content string     `json:"content" gorm:"text"`
	Status  NewsStatus `json:"status"`
	Topic   TopicList  `gorm:"many2many:news_topics;"`
}

// TableName table name
func (News) TableName() string {
	return "news"
}

type NewsList []News

func (n NewsList) ToMapByTitle() map[string]News {
	toMap := make(map[string]News, len(n))
	for _, v := range n {
		toMap[v.Title] = v
	}
	return toMap
}
