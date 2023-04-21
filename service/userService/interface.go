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
	Gender   *int    `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	Company  *string `json:"company" binding:"omitempty,max=20" label:"公司"` // 用户所在公司（非必填）
}

// UserLoginRequest 用于接收前端传递的创建新用户所需信息的 JSON 数据
type UserLoginRequest struct {
	Account  string `json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Password string `json:"password" binding:"required,min=6,max=40" label:"密码"` // 用户密码（未加密）
}

type UserRegisterRequest struct {
	Account   string         `json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Password  string         `json:"password" binding:"required,min=6,max=40" label:"密码"` // 用户密码（未加密）
	Name      string         `json:"name" binding:"required,min=2,max=20" label:"用户名"`
	Email     string         `json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile    string         `json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender    int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	Birthday  utils.DateTime `json:"birthday" binding:"omitempty"`
	Signature string         `json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	Company   string         `json:"company" binding:"omitempty,max=20" label:"公司"`
	Website   string         `json:"website" binding:"omitempty,max=255" label:"个人网站"`
}

func (registerData *UserRegisterRequest) ConvertToBBSUser(user *model.BBSUser) {
	user.Account = registerData.Account
	user.Password = registerData.Password
	user.Name = registerData.Name
	user.Email = registerData.Email
	user.Mobile = registerData.Mobile
	user.Gender = registerData.Gender
	user.Birthday = registerData.Birthday
	user.Signature = registerData.Signature
	user.Company = registerData.Company
	user.Website = registerData.Website
}

/*BaseBBSUserResponse
* @Description: 基本的用户信息
 */
type BaseBBSUserResponse struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id" label:"用户ID"`
	Account     string         `gorm:"type:varchar(20);not null;comment:登录账号，长度4~20" json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Name        string         `gorm:"type:varchar(20);not null;comment:用户名" json:"name" binding:"required,max=20" label:"用户名"`
	Status      int            `gorm:"type:tinyint(4);not null;default:0;comment:用户状态（10 正常、1 禁言、2 用户已离职、3 账号已删除、4 注册审核中）" json:"status" label:"用户状态"`
	StatusText  string         `json:"statusText" label:"用户状态"`
	Email       string         `gorm:"type:varchar(60);comment:邮箱" json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile      string         `gorm:"type:varchar(11);comment:手机号码" json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender      int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	GenderText  string         `json:"genderText" label:"性别"`
	Birthday    utils.DateTime `json:"birthday" label:"生日"`
	Signature   string         `gorm:"type:varchar(255);comment:个性签名" json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	IsAdmin     int            `gorm:"type:tinyint(4);default:0;comment:权限等级(0 普通、10 管理员、20 超级管理员）" json:"isAdmin" label:"权限等级"`
	IsAdminText string         `json:"isAdminText" label:"权限等级"`
	Company     string         `gorm:"type:varchar(20);comment:用户所在公司" json:"company" binding:"omitempty,max=20" label:"公司"`
	Website     string         `gorm:"type:varchar(255);comment:用户个人网站" json:"website" binding:"omitempty,max=255" label:"个人网站"`
	CreatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:创建时间" json:"createdAt" label:"创建时间"`
	UpdatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:更新时间" json:"updatedAt" label:"更新时间"`
}

/*AdvanceBBSUserResponse
* @Description: 进阶用户信息
 */
type AdvanceBBSUserResponse struct {
	// BaseBBSUserResponse 基础用户信息
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id" label:"用户ID"`
	Account     string         `gorm:"type:varchar(20);not null;comment:登录账号，长度4~20" json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Name        string         `gorm:"type:varchar(20);not null;comment:用户名" json:"name" binding:"required,max=20" label:"用户名"`
	Status      int            `gorm:"type:tinyint(4);not null;default:0;comment:用户状态（10 正常、1 禁言、2 用户已离职、3 账号已删除、4 注册审核中）" json:"status" label:"用户状态"`
	StatusText  string         `json:"statusText" label:"用户状态"`
	Email       string         `gorm:"type:varchar(60);comment:邮箱" json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile      string         `gorm:"type:varchar(11);comment:手机号码" json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender      int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	GenderText  string         `json:"genderText" label:"性别"`
	Birthday    utils.DateTime `json:"birthday" label:"生日"`
	Signature   string         `gorm:"type:varchar(255);comment:个性签名" json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	IsAdmin     int            `gorm:"type:tinyint(4);default:0;comment:权限等级(0 普通、10 管理员、20 超级管理员）" json:"isAdmin" label:"权限等级"`
	IsAdminText string         `json:"isAdminText" label:"权限等级"`
	Company     string         `gorm:"type:varchar(20);comment:用户所在公司" json:"company" binding:"omitempty,max=20" label:"公司"`
	Website     string         `gorm:"type:varchar(255);comment:用户个人网站" json:"website" binding:"omitempty,max=255" label:"个人网站"`
	CreatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:创建时间" json:"createdAt" label:"创建时间"`
	UpdatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:更新时间" json:"updatedAt" label:"更新时间"`

	// 额外用户信息
	LastLoginAt utils.DateTime `gorm:"type:datetime;comment:最后登录时间" json:"lastLoginAt" label:"最后登录时间"`
	LastLoginIP string         `gorm:"type:varchar(40);comment:最后登录IP地址，支持IPV6" json:"lastLoginIP" label:"最后登录IP地址"`
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

/*GetAllUsersInfoRequest
* @Description: 获取所有用户信息的请求体
 */
type GetAllUsersInfoRequest struct {
	Page uint `json:"page" binding:"required,min=1" label:"页码"`
}

/*GetAllUsersInfoResponse
* @Description: 获取所有用户信息的返回内容
 */
type GetAllUsersInfoResponse struct {
	Page         uint `json:"page" label:"页码"`
	TotalPage    uint `json:"total_page" label:"页码总数"`
	TotalUser    uint `json:"total_user" label:"用户总数"`
	CountPerPage uint `json:"count_per_page" label:"每页用户数量"`
	// 基础的用户信息
	List []BaseBBSUserResponse `json:"list"`
}

/*BatchUpdateUserStatusRequest
* @Description: 批量更新用户状态
 */
type BatchUpdateUserStatusRequest struct {
	List   []uint `json:"list" binding:"required,dive,required,min=1" label:"用户信息列表"`
	Status int    `json:"status" binding:"required,oneof=1 2 3 4 10" label:"状态"`
}

/*UpdateSelfInfoRequest
* @Description: 更新用户信息
 */
type UpdateSelfInfoRequest struct {
	Name      string         `json:"name" binding:"omitempty,min=2,max=20" label:"用户名"`
	Email     string         `json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile    string         `json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender    int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	Birthday  utils.DateTime `json:"birthday" binding:"omitempty"`
	Signature string         `json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	Company   string         `json:"company" binding:"omitempty,max=20" label:"公司"`
	Website   string         `json:"website" binding:"omitempty,max=255" label:"个人网站"`
}

/*UpdateUserInfoRequest
* @Description: 更新指定用户信息
 */
type UpdateUserInfoRequest struct {
	ID        uint           `json:"id" binding:"required" label:"用户ID"`
	Name      string         `json:"name" binding:"omitempty,min=2,max=20" label:"用户名"`
	Email     string         `json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile    string         `json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender    int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	Birthday  utils.DateTime `json:"birthday" binding:"omitempty"`
	Signature string         `json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	Company   string         `json:"company" binding:"omitempty,max=20" label:"公司"`
	Website   string         `json:"website" binding:"omitempty,max=255" label:"个人网站"`
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
	baseBBSUser.Status = user.Status
	baseBBSUser.StatusText = user.GetBBSStatusText()
	baseBBSUser.Email = user.Email
	baseBBSUser.Mobile = user.Mobile
	baseBBSUser.Gender = user.Gender
	baseBBSUser.GenderText = user.GetGender()
	baseBBSUser.Birthday = user.Birthday
	baseBBSUser.Signature = user.Signature
	baseBBSUser.IsAdmin = user.IsAdmin
	baseBBSUser.IsAdminText = user.GetAdminStatus()
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
	advanceBBSUser.Status = user.Status
	advanceBBSUser.StatusText = user.GetBBSStatusText()
	advanceBBSUser.Email = user.Email
	advanceBBSUser.Mobile = user.Mobile
	advanceBBSUser.Gender = user.Gender
	advanceBBSUser.GenderText = user.GetGender()
	advanceBBSUser.Birthday = user.Birthday
	advanceBBSUser.Signature = user.Signature
	advanceBBSUser.IsAdmin = user.IsAdmin
	advanceBBSUser.IsAdminText = user.GetAdminStatus()
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
