package model

import "main/utils"

/*Post
* @Description: 回复贴的表
 */
type Post struct {
	ID           uint           `gorm:"primaryKey;autoIncrement;comment:帖子ID" json:"id" label:"帖子ID"`
	Content      string         `gorm:"type:text;size:4000;comment:帖子内容，不超过4000个字符" json:"content" binding:"required,max=4000" label:"帖子内容"`
	IsDel        bool           `gorm:"comment:帖子状态（false正常，true已删除）" json:"isDel" label:"是否已删除"`
	CreateIP     string         `gorm:"type:varchar(39);comment:发帖的IP地址，支持IPV6" json:"createIP" label:"发帖的IP地址"`
	CreatedAt    utils.DateTime `gorm:"type:datetime;not null;comment:发帖时间" json:"createdAt" label:"发帖时间"`
	LastReplyAt  utils.DateTime `gorm:"type:datetime;not null;comment:最后回复时间" json:"lastReplyAt" label:"最后回复时间"`
	CreateUserID uint           `gorm:"comment:发帖人的ID" json:"createUserID" label:"发帖人的ID"`
	TopicID      uint           `gorm:"comment:主题帖的ID" json:"topicID" label:"主题帖的ID"`
}

func (post Post) TableName() string {
	return "post"
}
