package topicService

import "main/utils"

/*CreateTopicRequest
* @Description: 创建主题帖的请求
 */
type CreateTopicRequest struct {
	Title   string `gorm:"type:varchar(30);comment:帖子标题，不超过30个字" json:"title" binding:"required,max=30" label:"帖子标题"`
	Content string `gorm:"type:text;size:4000;comment:内容，不超过4000个字符" json:"content" binding:"required,max=4000" label:"内容"`
}

/*GetTopicListRequest
* @Description: 获取所有帖子的请求体
 */
type GetTopicListRequest struct {
	Page uint `json:"page" binding:"required,min=1" label:"页码"`
	// 查看被删除的帖子，为另一个接口，这里不增加复杂度
	//IsHasDel bool `json:"IsHasDel" binding:"omitempty" label:"是否包含已删除，只有管理员才能设置该参数为true"`
}

/*GetNormalTopicResponse
* @Description: 获取所有帖子的返回内容
 */
type GetNormalTopicResponse struct {
	Page         uint `json:"page" label:"页码"`
	TotalPage    uint `json:"total_page" label:"页码总数"`
	TotalTopic   uint `json:"total_topic" label:"帖子总数"`
	CountPerPage uint `json:"count_per_page" label:"每页用户数量"`
	// 基础的用户信息
	List []AllTopicResponse `json:"list"`
}

/*AllTopicResponse
* @Description: 主题帖的信息
 */
type AllTopicResponse struct {
	ID             uint           `gorm:"primaryKey;autoIncrement;comment:帖子ID" json:"id" label:"帖子ID"`
	Content        string         `gorm:"type:text;size:4000;comment:内容，不超过4000个字符" json:"content" binding:"required,max=4000" label:"内容"`
	CreateIP       string         `gorm:"type:varchar(39);comment:发帖的IP地址，支持IPV6" json:"createIP" label:"发帖的IP地址"`
	CreatedAt      utils.DateTime `gorm:"type:datetime;not null;comment:发帖时间" json:"createdAt" label:"发帖时间"`
	CreateUserID   uint           `gorm:"comment:发帖人的ID" json:"createUserID" label:"发帖人的ID"`
	CreateUserName string         `json:"createUserName" label:"发帖人姓名"`
	LastReplyAt    utils.DateTime `gorm:"type:datetime;not null;comment:最后回复时间" json:"last_reply_at" label:"更新时间"`
	Title          string         `gorm:"type:varchar(30);comment:帖子标题，不超过30个字" json:"title" binding:"required,max=30" label:"帖子标题"`
	ReplayCount    uint           `json:"replay_count" label:"回复数量"`
}
