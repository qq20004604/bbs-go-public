package topicService

import (
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/db"
	"main/db/model"
	"main/service/userService"
)

/*CreateTopic
* @Description: 创建一个主题帖
* @param data
* @param c
* @return error
 */
func CreateTopic(data CreateTopicRequest, userID uint, c *gin.Context) error {
	// 1. 发帖检查
	if err := IsCanCreateTopic(userID); err != nil {
		return err
	}

	var topic = model.Topic{
		CreateUserID: userID,
		Content:      data.Content,
		Title:        data.Title,
	}
	ip := userService.GetClientIP(c)
	topic.BeforeCreate(ip)
	if err := db.DbItem.Model(&model.Topic{}).Create(&topic).Error; err != nil {
		log.Error(err)
		return errors.New("发帖失败，请重试或者联系管理员")
	} else {
		err2 := AfterSuccessCreateTopic(userID)
		if err2 != nil {
			log.Error(err2)
		}
		return nil
	}
}

/*IsCanCreateTopic
* @Description: 是否可以创建主题
* @param userID
* @return error
 */
func IsCanCreateTopic(userID uint) error {
	// 2.1 发帖成功后，5分钟后才能进行下一次发帖。（redis加300秒失效的key来实现）（非成功的，通过限流器来设置）
	// todo 懒得写了

	// 2.2 发帖人状态检查（只有 status=正常的人才能发帖）
	var user model.BBSUser
	err := db.DbItem.Model(&model.BBSUser{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		log.Error("用户查询失败，id=", userID)
		return errors.New("服务器错误")
	} else if user.ID == 0 {
		// 其实理论上是不可能进入这个分支，因为在调用这个函数前，会检查用户是否在线，在那一步已经查询了该用户是否存在了
		return errors.New("该用户不存在，无法发帖")
	} else if user.Status != model.UserStatusNormal {
		return errors.New("该用户已被禁止发帖")
	}

	// 2.3 其他检查

	return nil
}

/*AfterSuccessCreateTopic
* @Description: 在用户成功发帖之后
* @param userID
* @return error
 */
func AfterSuccessCreateTopic(userID uint) error {
	// 1. 配合发帖检查使用，发帖成功后在redis里写一个5分钟失效的缓存

	// 2. 其他处理

	return nil
}
