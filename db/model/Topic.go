package model

import (
	"gorm.io/gorm"
	"main/utils"
	"time"
)

/*Topic
* @Description: 主题帖
 */
type Topic struct {
	ID           uint           `gorm:"primaryKey;autoIncrement;comment:帖子ID" json:"id" label:"帖子ID"`
	Content      string         `gorm:"type:text;size:4000;comment:内容，不超过4000个字符" json:"content" binding:"required,max=4000" label:"内容"`
	IsDel        bool           `gorm:"comment:帖子状态（false正常，true已删除）" json:"isDel" label:"是否已删除"`
	CreateIP     string         `gorm:"type:varchar(39);comment:发帖的IP地址，支持IPV6" json:"createIP" label:"发帖的IP地址"`
	CreatedAt    utils.DateTime `gorm:"type:datetime;not null;comment:发帖时间" json:"createdAt" label:"发帖时间"`
	CreateUserID uint           `gorm:"comment:发帖人的ID" json:"createUserID" label:"发帖人的ID"`
	LastReplyAt  utils.DateTime `gorm:"type:datetime;not null;comment:最后回复时间" json:"last_reply_at" label:"更新时间"`
	Title        string         `gorm:"type:varchar(30);comment:帖子标题，不超过30个字" json:"title" binding:"required,max=30" label:"帖子标题"`
}

func (topic Topic) TableName() string {
	return "topic"
}

/*AfterReply
* @Description: 有人回帖后，调用本方法
 */
func (topic *Topic) AfterReply(tx *gorm.DB) (err error) {
	topic.LastReplyAt = utils.DateTime{Time: time.Now()}
	return
}

func (topic *Topic) BeforeCreate(ip string) {
	topic.CreateIP = ip
	topic.CreatedAt = utils.DateTime{Time: time.Now()}
	topic.LastReplyAt = topic.CreatedAt
	return
}
