package topicService

type CreateTopicRequest struct {
	Title   string `gorm:"type:varchar(30);comment:帖子标题，不超过30个字" json:"title" binding:"required,max=30" label:"帖子标题"`
	Content string `gorm:"type:text;size:4000;comment:内容，不超过4000个字符" json:"content" binding:"required,max=4000" label:"内容"`
}
