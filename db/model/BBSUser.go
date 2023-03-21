package model

import (
	"main/utils"
)

/*BBSUser
* @Description: 论坛用户表结构设计
 */
type BBSUser struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;comment:用户ID" json:"id"`
	AuthID      uint           `gorm:"type:int;not null;comment:在Auth系统里的ID，识别用户时，以这个为准" json:"auth_id"`
	Account     string         `gorm:"type:varchar(20);not null;comment:登录账号" json:"account"`
	Name        string         `gorm:"type:varchar(20);not null;comment:用户名" json:"name"`
	Password    string         `gorm:"type:varchar(40);not null;comment:密码" json:"password"`
	Status      int            `gorm:"type:tinyint(4);not null;default:0;comment:用户状态（0 正常、1 禁言、2 用户已离职、3 账号已删除）" json:"status"`
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
	CreatedAt   utils.DateTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt   utils.DateTime `gorm:"type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

func (v BBSUser) TableName() string {
	return "bbs_user"
}
