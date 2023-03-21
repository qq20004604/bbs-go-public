package userService

import (
	"main/db/model"
	"main/utils"
)

// CreateUserRequest 用于接收前端传递的创建新用户所需信息的 JSON 数据
type CreateUserRequest struct {
	Account  string  `json:"account"`
	Name     string  `json:"name" binding:"required,min=2,max=20" label:"用户名"`    // 用户名
	Password string  `json:"password" binding:"required,min=6,max=40" label:"密码"` // 用户密码（未加密）
	Email    *string `json:"email" binding:"omitempty,email" label:"邮箱"`          // 用户邮箱（非必填）
	Mobile   *string `json:"mobile" binding:"omitempty,len=11" label:"手机号"`       // 用户手机号码（非必填）
	Gender   *int    `json:"gender" binding:"omitempty,oneof=0 1 2" label:"性别"`   // 用户性别（0 未知、1 男、2 女，非必填）
	Company  *string `json:"company" binding:"omitempty,max=20" label:"公司"`       // 用户所在公司（非必填）
}

/*BaseUserInfoResponse
* @Description: 基本的用户信息
 */
type BaseUserInfoResponse struct {
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

/*AdvanceUserInfoResponse
* @Description: 进阶用户信息
 */
type AdvanceUserInfoResponse struct {
	BaseUserInfoResponse

	AuthID      uint           `json:"auth_id"`
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

/*getStatusText
* @Description: 获取用户状态的文字信息
* @param status
* @return string
 */
func getStatusText(status int) string {
	switch status {
	case 0:
		return "正常"
	case 1:
		return "禁言"
	case 2:
		return "用户已离职"
	case 3:
		return "账号已删除"
	default:
		return "未知状态"
	}
}

/*getGender
* @Description: 获取性别
* @param status
* @return string
 */
func getGender(status int) string {
	switch status {
	case 1:
		return "男"
	case 2:
		return "女"
	case 3:
		return "男同"
	case 4:
		return "女同"
	default:
		return "未知状态"
	}
}

/*getAdminStatus
* @Description: 获取管理员相关状态
* @param status
* @return string
 */
func getAdminStatus(status int) string {
	switch status {
	case 10:
		return "管理员"
	case 20:
		return "超级管理员"
	default:
		return "普通"
	}
}

/*ConvertBBSUserToBaseUserInfoResponse
* @Description: 将数据库模型 BBSUser 的数据转为 UserInfoResponse 的结构
* @param userService
* @return UserInfoResponse
 */
func ConvertBBSUserToBaseUserInfoResponse(user *model.BBSUser) BaseUserInfoResponse {
	var userRes = BaseUserInfoResponse{
		ID:        user.ID,
		Account:   user.Account,
		Name:      user.Name,
		Status:    getStatusText(user.Status),
		Email:     user.Email,
		Mobile:    user.Mobile,
		Gender:    getGender(user.Gender),
		Birthday:  utils.DateTime{},
		Signature: user.Signature,
		IsAdmin:   getAdminStatus(user.IsAdmin),
		Company:   user.Company,
		Website:   user.Website,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return userRes
}

/*ConvertBBSUserToAdvanceUserInfoResponse
* @Description: 返回更进一步的用户信息
* @param userService
* @return AdvanceUserInfoResponse
 */
func ConvertBBSUserToAdvanceUserInfoResponse(user *model.BBSUser) AdvanceUserInfoResponse {
	var userRes = AdvanceUserInfoResponse{
		BaseUserInfoResponse: BaseUserInfoResponse{
			ID:        user.ID,
			Account:   user.Account,
			Name:      user.Name,
			Status:    getStatusText(user.Status),
			Email:     user.Email,
			Mobile:    user.Mobile,
			Gender:    getGender(user.Gender),
			Birthday:  utils.DateTime{},
			Signature: user.Signature,
			IsAdmin:   getAdminStatus(user.IsAdmin),
			Company:   user.Company,
			Website:   user.Website,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		//ID:          userService.ID,
		//Account:     userService.Account,
		//Name:        userService.Name,
		//Status:      getStatusText(userService.Status),
		//Email:       userService.Email,
		//Mobile:      userService.Mobile,
		//Gender:      getGender(userService.Gender),
		//Birthday:    utils.DateTime{},
		//Signature:   userService.Signature,
		//IsAdmin:     getAdminStatus(userService.IsAdmin),
		//Company:     userService.Company,
		//Website:     userService.Website,
		//CreatedAt:   userService.CreatedAt,
		//UpdatedAt:   userService.UpdatedAt,
		AuthID:      user.AuthID,
		LastLoginAt: user.LastLoginAt,
		LastLoginIP: user.LastLoginIP,
	}

	return userRes
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
