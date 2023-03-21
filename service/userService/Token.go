package userService

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db"
)

/*MakeToken
* @Description: 随机生成 token
* @return string	生成的token字符串，其长度 = 20 + 4
* @return error
 */
func MakeToken() (string, error) {
	const Length = 20
	bytes := make([]byte, Length/2)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Error("函数 MakeToken 生成随机字符串错误：" + err.Error())
		return "", err
	}
	return "bbs-" + hex.EncodeToString(bytes), nil
}

/*GetToken
* @Description: 获取当前用户的token
* @param c
* @return string
* @return error
 */
func GetToken(c *gin.Context) (string, error) {
	// 1、先拿到用户 token
	token := c.Request.Header.Get(config.Config.Common.HeaderTokenName)

	// 2. 如果 token 为空，返回错误
	if token == "" {
		return "", errors.New("登录过期，请登录")
	}
	return token, nil
}

/*IsTokenExist
* @Description: 根据入参，去redis里查询该token是否存在
* @param token
 */
func IsTokenExist(c *gin.Context, token string) (bool, error) {
	if res, err := db.RedisDB.Exists(c, token).Result(); err != nil {
		return false, err
	} else if res > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

/*GetUserInfoFromToken
* @Description: 拿token，去redis里获取用户信息
* @param c
 */
func GetUserInfoFromToken(c *gin.Context) error {
	token, errGetToken := GetToken(c)
	if errGetToken != nil {
		return errGetToken
	}

	// 3. 连接 redis，根据 redis key，获取对应的value
	redisKey := "bbs-" + token
	println(redisKey)

	return nil
}
