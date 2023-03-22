package userService

import (
	"crypto/sha256"
	"encoding/hex"
	"main/config"
)

/*SaltAndHashPassword
* @Description:  对密码进行加盐处理，使用 SHA-256 进行加盐
* @param account
* @param password
* @return string
 */
func SaltAndHashPassword(account, password string) string {
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
