package topicService

import (
	log "github.com/sirupsen/logrus"
	"main/db"
	"main/db/model"
)

/*GetNormalTopicListByPage
* @Description: 返回对应页码的主题数据（限定：未被删除）
* @param page		页码
* @return GetAllUsersInfoResponse	对应页码的数据
 */
func GetNormalTopicListByPage(page uint) GetNormalTopicResponse {
	// 每页主题帖数目
	const TopicPerPage = 20

	// 1. 先获取帖子总数，根据单页数量，计算总共有多少页数据
	var totalTopic int64
	db.DbItem.Model(model.Topic{}).Where("is_del = ?", false).Count(&totalTopic)
	totalPages := (totalTopic + TopicPerPage - 1) / TopicPerPage

	// 如果请求页码的数据大于当前页码
	if page > uint(totalPages) {
		return GetNormalTopicResponse{
			Page:         page,
			TotalPage:    uint(totalPages),
			TotalTopic:   uint(totalTopic),
			CountPerPage: TopicPerPage,
			List:         make([]AllTopicResponse, 0),
		}
	}

	// 2. 根据当前页码，返回对应页码的用户列表
	offset := (page - 1) * TopicPerPage
	var resList []AllTopicResponse
	// todo 优化这里的查询数据
	err := db.DbItem.Table("topic").
		Select("topic.*, bbs_user.name as create_user_name").
		Joins("LEFT JOIN bbs_user ON bbs_user.id = topic.create_user_id").
		Where("topic.is_del = ?", false).
		Order("topic.last_reply_at DESC").
		Offset(int(offset)).
		Limit(TopicPerPage).
		Find(&resList).Error
	if err != nil {
		log.Error(err)
		return GetNormalTopicResponse{
			Page:         page,
			TotalPage:    uint(totalPages),
			TotalTopic:   uint(totalTopic),
			CountPerPage: TopicPerPage,
			List:         make([]AllTopicResponse, 0),
		}
	}

	return GetNormalTopicResponse{
		Page:         page,
		TotalPage:    uint(totalPages),
		TotalTopic:   uint(totalTopic),
		CountPerPage: TopicPerPage,
		List:         resList,
	}
}
