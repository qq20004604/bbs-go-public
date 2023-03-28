package model

import (
	"gorm.io/gorm"
	"main/utils"
	"time"
)

/*BBSUser
* @Description: 论坛用户表结构设计
 */
type BBSUser struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id" label:"用户ID"`
	AuthID      uint           `gorm:"type:int;not null;comment:在Auth系统里的ID，识别用户时，以这个为准" json:"authID"`
	Account     string         `gorm:"type:varchar(20);not null;comment:登录账号，长度4~20" json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Name        string         `gorm:"type:varchar(20);not null;comment:用户名" json:"name" binding:"required,max=20" label:"用户名"`
	Password    string         `gorm:"type:varchar(40);not null;comment:密码" json:"password" binding:"required,min=6,max=40" label:"密码"`
	Status      int            `gorm:"type:tinyint(4);not null;default:0;comment:用户状态（0 正常、1 禁言、2 用户已离职、3 账号已删除、4 注册审核中）" json:"status" label:"用户状态"`
	LastLoginAt utils.DateTime `gorm:"type:datetime;comment:最后登录时间" json:"lastLoginAt" label:"最后登录时间"`
	LastLoginIP string         `gorm:"type:varchar(40);comment:最后登录IP地址，支持IPV6" json:"lastLoginIP" label:"最后登录IP地址"`
	Email       string         `gorm:"type:varchar(60);comment:邮箱（长度最大为60）" json:"email" binding:"omitempty,email,max=60" label:"邮箱"`
	Mobile      string         `gorm:"type:varchar(11);comment:手机号码" json:"mobile" binding:"omitempty,len=11" label:"手机号码"`
	Gender      int            `gorm:"type:tinyint(4);default:0;comment:性别（1 男、2 女、3 男同、4 女同、5 未告知、6 其他）" json:"gender" binding:"omitempty,gte=1,lte=6" label:"性别"`
	Birthday    utils.DateTime `gorm:"type:date;comment:生日" json:"birthday" label:"生日"`
	Signature   string         `gorm:"type:varchar(255);comment:个性签名" json:"signature" binding:"omitempty,max=255" label:"个性签名"`
	IsAdmin     int            `gorm:"type:tinyint(4);default:0;comment:权限等级(0 普通、10 管理员、20 超级管理员）" json:"isAdmin" label:"权限等级"`
	Company     string         `gorm:"type:varchar(20);comment:公司" json:"company" binding:"omitempty,max=20" label:"公司"`
	Website     string         `gorm:"type:varchar(255);comment:个人网站" json:"website" binding:"omitempty,max=255" label:"个人网站"`
	CreatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:创建时间" json:"createdAt" label:"创建时间"`
	UpdatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:更新时间" json:"updatedAt" label:"更新时间"`
}

func (user BBSUser) TableName() string {
	return "bbs_user"
}

func (user *BBSUser) BeforeUpdate(tx *gorm.DB) (err error) {
	user.UpdatedAt = utils.DateTime{Time: time.Now()}
	return
}

func (user *BBSUser) UpdateAfterLogin(ip string) {
	user.LastLoginIP = ip
	user.LastLoginAt = utils.DateTime{Time: time.Now()}
	return
}

const (
	UserStatusNormal        = 0
	UserStatusMuted         = 1
	UserStatusResigned      = 2
	UserStatusDeleted       = 3
	UserStatusPendingReview = 4
)

/*IsUserStatusNormal
* @Description: 当前用户是否正常
* @return string
 */
func (user *BBSUser) IsUserStatusNormal() bool {
	if user.Status == UserStatusNormal {
		return true
	} else {
		return false
	}
}

/*GetBBSStatusText
* @Description: 获取用户状态的文字信息
* @return string
 */
func (user *BBSUser) GetBBSStatusText() string {
	switch user.Status {
	case UserStatusNormal:
		return "正常"
	case UserStatusMuted:
		return "禁言"
	case UserStatusResigned:
		return "用户已离职"
	case UserStatusDeleted:
		return "账号已删除"
	case UserStatusPendingReview:
		return "注册审核中"
	default:
		return "未知状态"
	}
}

/*GetGender
* @Description: 获取性别的文字信息
* @return string
 */
func (user *BBSUser) GetGender() string {
	switch user.Gender {
	case 1:
		return "男"
	case 2:
		return "女"
	case 3:
		return "男同"
	case 4:
		return "女同"
	case 5:
		return "未告知"
	case 6:
		return "其他"
	default:
		return "未知"
	}
}

/*GetAdminStatus
* @Description: 获取管理员相关状态
* @param status
* @return string
 */
func (user *BBSUser) GetAdminStatus() string {
	switch user.IsAdmin {
	case 10:
		return "管理员"
	case 20:
		return "超级管理员"
	default:
		return "普通"
	}
}
