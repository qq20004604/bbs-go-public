package userService

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db"
	"main/db/model"
)

/*CreateUser
* @Description:	向数据库中插入一条新的用户记录
* @param userService
* @return error
 */
func CreateUser(user *model.BBSUser) error {
	// 再基于用户数据并插入到数据库表之前，先执行一些前置操作，例如密码加盐等
	if err := beforeCreateUser(user); err != nil {
		log.Error(err)
		return err
	}

	// 执行插入操作
	result := db.DbItem.Model(model.BBSUser{}).Create(user)

	// 插入失败
	if result.Error != nil {
		//log.WithFields(log.Fields{
		//	"userService": user.Name,
		//	"error":       result.Error,
		//}).Error("Failed to create userService")
		return result.Error
	}

	// 插入成功
	//log.WithField("userService", user.Name).Info("User created successfully")
	return nil
}

// saltAndHashPassword 通过给定的账号和密码生成一个哈希后的密码字符串。
// 它使用配置文件中的预定义盐值对密码进行加盐，然后将其与账号进行拼接。
// 最后，使用 SHA-256 算法计算拼接后字符串的哈希值，并将其转换为十六进制字符串。
// 返回截取到指定长度的哈希后的密码字符串。
func saltAndHashPassword(account, password string) string {
	// 从配置文件中获取预定义的盐值
	salt := config.Config.Common.PasswordSalt
	// 拼接密码、盐值和账号
	saltedPassword := password + salt + account

	// 使用 SHA-256 哈希算法计算拼接后字符串的哈希值
	hashedPassword := sha256.Sum256([]byte(saltedPassword))
	// 将字节数组的哈希值转换为十六进制字符串
	hashedPasswordStrFull := hex.EncodeToString(hashedPassword[:])

	// 从配置文件中获取哈希后密码字符串的指定长度
	passwordStrLength := config.Config.Common.PasswordLengthAfterHash
	if passwordStrLength > 64 {
		passwordStrLength = 64
	} else if passwordStrLength < 8 {
		passwordStrLength = 8
	}
	// 截取指定长度的哈希后的密码字符串
	hashedPasswordStr := hashedPasswordStrFull[:passwordStrLength]

	return hashedPasswordStr
}

/*beforeCreateUser
* @Description: 在创建用户之前，进行一些特殊处理
* @param user	用户信息
* @return error
 */
func beforeCreateUser(user *model.BBSUser) error {
	// 1. 对密码进行加密处理，只存储加密后的密码
	pwWithSalt := saltAndHashPassword(user.Name, user.Password)
	user.Password = pwWithSalt

	// 2. 对用户进行查重，如果该用户重复的话，则不允许创建
	var userData = BBSUserExist{
		AuthID:  user.AuthID,
		Account: user.Account,
		Name:    user.Name,
		Email:   user.Email,
		Mobile:  user.Mobile,
	}
	if isExist, err := IsUserExistByAccount(&userData); err != nil {
		return err
	} else if isExist {
		return errors.New("该用户已存在")
	}

	// final. 一切正常
	return nil
}
