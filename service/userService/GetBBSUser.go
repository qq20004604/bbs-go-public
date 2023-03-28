package userService

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/db"
	"main/db/model"
)

/*GetAdvanceBBSUserInfoBySelf
* @Description: 查询当前用户的个人信息
* @param c
* @param ID
* @return error
 */
func GetAdvanceBBSUserInfoBySelf(c *gin.Context) (error, AdvanceBBSUserResponse) {
	// 1. 在线检查
	isOnline, userID := IsOnline(c)
	if !isOnline {
		return errors.New("未登录，请登录后再试"), AdvanceBBSUserResponse{}
	}

	// 2. 查询用户数据
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("id = ?", userID).First(&user)
	if user.ID == 0 {
		return errors.New("该用户不存在"), AdvanceBBSUserResponse{}
	}

	// 3. 检查用户状态，如果不正常，则返回其状态。如果禁止某个状态用户登录，则修改这个状态
	// 但理论上，能登录，就意味着这个账号能正常获取自己的个人数据
	if user.IsUserStatusNormal() == false {
		return errors.New(user.GetBBSStatusText()), AdvanceBBSUserResponse{}
	}

	// 4. 将数据转为返回给用户的格式
	var resUserData AdvanceBBSUserResponse
	resUserData.ConvertFromBBSUser(&user)

	// 5. 返回数据给调用方
	return nil, resUserData
}

/*GetBBSUserByAccount
* @Description: 根据用户账号获取用户信息
* @param c
* @param account	用户账号
* @return error
* @return model.BBSUser
 */
func GetBBSUserByAccount(account string) (error, model.BBSUser) {
	var user model.BBSUser
	db.DbItem.Model(model.BBSUser{}).Where("account = ?", account).First(&user)
	if user.ID == 0 {
		return errors.New("该用户不存在"), model.BBSUser{}
	}
	return nil, user
}
