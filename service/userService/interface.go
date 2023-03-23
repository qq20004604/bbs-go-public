package userService

import (
	"main/db/model"
	"main/utils"
)

// CreateUserRequest 用于接收前端传递的创建新用户所需信息的 JSON 数据
type CreateUserRequest struct {
	Account  string  `json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Name     string  `json:"name" binding:"required,min=2,max=20" label:"用户名"`    // 用户名
	Password string  `json:"password" binding:"required,min=6,max=40" label:"密码"` // 用户密码（未加密）
	Email    *string `json:"email" binding:"omitempty,email" label:"邮箱"`          // 用户邮箱（非必填）
	Mobile   *string `json:"mobile" binding:"omitempty,len=11" label:"手机号"`       // 用户手机号码（非必填）
	Gender   *int    `json:"gender" binding:"omitempty,oneof=0 1 2" label:"性别"`   // 用户性别（0 未知、1 男、2 女，非必填）
	Company  *string `json:"company" binding:"omitempty,max=20" label:"公司"`       // 用户所在公司（非必填）
}

// UserLoginRequest 用于接收前端传递的创建新用户所需信息的 JSON 数据
type UserLoginRequest struct {
	Account  string `json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Password string `json:"password" binding:"required,min=6,max=40" label:"密码"` // 用户密码（未加密）
}

/*BaseBBSUserResponse
* @Description: 基本的用户信息
 */
type BaseBBSUserResponse struct {
	ID        uint           `json:"id"`
	Account   string         `json:"account"`
	Name      string         `json:"name"`
	Status    string         `json:"status"`
	Email     string         `json:"email"`
	Mobile    string         `json:"mobile"`
	Gender    string         `json:"gender"`
	Birthday  utils.DateTime `json:"birthday"`
	Signature string         `json:"signature"`
	IsAdmin   string         `json:"is_admin"`
	Company   string         `json:"company"`
	Website   string         `json:"website"`
	CreatedAt utils.DateTime `json:"created_at"`
	UpdatedAt utils.DateTime `json:"updated_at"`
}

/*AdvanceBBSUserResponse
* @Description: 进阶用户信息
 */
type AdvanceBBSUserResponse struct {
	// BaseBBSUserResponse 基础用户信息
	ID        uint           `json:"id"`
	Account   string         `json:"account"`
	Name      string         `json:"name"`
	Status    string         `json:"status"`
	Email     string         `json:"email"`
	Mobile    string         `json:"mobile"`
	Gender    string         `json:"gender"`
	Birthday  utils.DateTime `json:"birthday"`
	Signature string         `json:"signature"`
	IsAdmin   string         `json:"is_admin"`
	Company   string         `json:"company"`
	Website   string         `json:"website"`
	CreatedAt utils.DateTime `json:"created_at"`
	UpdatedAt utils.DateTime `json:"updated_at"`

	// 额外用户信息
	//AuthID      uint           `json:"auth_id"`
	LastLoginAt utils.DateTime `json:"last_login_at"`
	LastLoginIP string         `json:"last_login_ip"`
}

/*TokenUserInfo
* @Description: 存储在 token 里的用户信息，只存最基本的信息，需要获取额外信息时则从数据库里读取
 */
type TokenUserInfo struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	IsAdmin int    `json:"is_admin"`
}

/*BBSUserExist
* @Description: 用于判断用户是否存在的结构体，本结构体里的字段，在数据库里应该都是唯一的
 */
type BBSUserExist struct {
	AuthID  uint   `json:"auth_id"`
	Account string `json:"account"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
}

/*ConvertFromBBSUser
* @Description: 将数据库模型 BBSUser 的数据转为 UserInfoResponse 的结构
* @param userService
* @return UserInfoResponse
 */
func (baseBBSUser *BaseBBSUserResponse) ConvertFromBBSUser(user *model.BBSUser) {
	baseBBSUser.ID = user.ID
	baseBBSUser.Account = user.Account
	baseBBSUser.Name = user.Name
	baseBBSUser.Status = user.GetBBSStatusText()
	baseBBSUser.Email = user.Email
	baseBBSUser.Mobile = user.Mobile
	baseBBSUser.Gender = user.GetGender()
	baseBBSUser.Birthday = utils.DateTime{}
	baseBBSUser.Signature = user.Signature
	baseBBSUser.IsAdmin = user.GetAdminStatus()
	baseBBSUser.Company = user.Company
	baseBBSUser.Website = user.Website
	baseBBSUser.CreatedAt = user.CreatedAt
	baseBBSUser.UpdatedAt = user.UpdatedAt
}

/*ConvertFromBBSUser
* @Description: 返回更进一步的用户信息
* @param userService
* @return AdvanceBBSUserResponse
 */
func (advanceBBSUser *AdvanceBBSUserResponse) ConvertFromBBSUser(user *model.BBSUser) {
	advanceBBSUser.ID = user.ID
	advanceBBSUser.Account = user.Account
	advanceBBSUser.Name = user.Name
	advanceBBSUser.Status = user.GetBBSStatusText()
	advanceBBSUser.Email = user.Email
	advanceBBSUser.Mobile = user.Mobile
	advanceBBSUser.Gender = user.GetGender()
	advanceBBSUser.Birthday = utils.DateTime{}
	advanceBBSUser.Signature = user.Signature
	advanceBBSUser.IsAdmin = user.GetAdminStatus()
	advanceBBSUser.Company = user.Company
	advanceBBSUser.Website = user.Website
	advanceBBSUser.CreatedAt = user.CreatedAt
	advanceBBSUser.UpdatedAt = user.UpdatedAt

	// 以上是基本用户信息，以下的 advanceBBSUser 里额外的信息
	//advanceBBSUser.AuthID = user.AuthID
	advanceBBSUser.LastLoginAt = user.LastLoginAt
	advanceBBSUser.LastLoginIP = user.LastLoginIP
}

/*ConvertBBSUserToTokenUserInfo
* @Description: 将用户数据转为用于写入redis里的数据格式
* @param userService
* @return TokenUserInfo
 */
func ConvertBBSUserToTokenUserInfo(user *model.BBSUser) TokenUserInfo {
	var data TokenUserInfo
	utils.MergeData(&data, user)
	return data
}
