package model

// Topic represent entity of the topic
type Topic struct {
	BaseModel
	Name string   `json:"name"`
	Slug string   `json:"slug"`
	News NewsList `gorm:"many2many:news_topics;" json:"news"`
}

// TableName table name
func (Topic) TableName() string {
	return "topics"
}

type TopicList []Topic

func (n TopicList) ToMapByTitle() map[string]Topic {
	toMap := make(map[string]Topic, len(n))
	for _, v := range n {
		toMap[v.Name] = v
	}
	return toMap
}
