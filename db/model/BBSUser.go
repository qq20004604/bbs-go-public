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
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	AuthID      uint           `gorm:"type:int;not null;comment:在Auth系统里的ID，识别用户时，以这个为准" json:"auth_id"`
	Account     string         `gorm:"type:varchar(20);not null;comment:登录账号，长度4~20" json:"account" binding:"required,min=4,max=20" label:"登录账号"`
	Name        string         `gorm:"type:varchar(20);not null;comment:用户名" json:"name"`
	Password    string         `gorm:"type:varchar(40);not null;comment:密码" json:"password" binding:"required,min=6,max=40" label:"密码"`
	Status      int            `gorm:"type:tinyint(4);not null;default:0;comment:用户状态（0 正常、1 禁言、2 用户已离职、3 账号已删除、4 注册审核中）" json:"status"`
	LastLoginAt utils.DateTime `gorm:"type:datetime;comment:用户最后登录时间" json:"last_login_at"`
	LastLoginIP string         `gorm:"type:varchar(40);comment:用户最后登录 IP 地址，支持IPV6" json:"last_login_ip"`
	Email       string         `gorm:"type:varchar(60);comment:用户邮箱（长度最大为60）" json:"email"`
	Mobile      string         `gorm:"type:varchar(11);comment:用户手机号码" json:"mobile"`
	Gender      int            `gorm:"type:tinyint(4);default:0;comment:用户性别（0 未知、1 男、2 女）" json:"gender"`
	Birthday    utils.DateTime `gorm:"type:date;comment:用户生日" json:"birthday"`
	Signature   string         `gorm:"type:varchar(255);comment:用户个性签名" json:"signature"`
	IsAdmin     int            `gorm:"type:tinyint(4);default:0;comment:是否为管理员(0 普通、10 管理员、20 超级管理员）" json:"is_admin"`
	Company     string         `gorm:"type:varchar(20);comment:用户所在公司" json:"company"`
	Website     string         `gorm:"type:varchar(255);comment:用户个人网站" json:"website"`
	CreatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:创建时间" json:"created_at"`
	UpdatedAt   utils.DateTime `gorm:"type:datetime;not null;comment:更新时间" json:"updated_at"`
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

/*GetBBSStatusText
* @Description: 获取用户状态的文字信息
* @return string
 */
func (user *BBSUser) GetBBSStatusText() string {
	switch user.Status {
	case 0:
		return "正常"
	case 1:
		return "禁言"
	case 2:
		return "用户已离职"
	case 3:
		return "账号已删除"
	case 4:
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
	default:
		return "未知性别"
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
