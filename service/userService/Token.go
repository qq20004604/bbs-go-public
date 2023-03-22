package userService

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db"
	"time"
)

/*MakeToken
* @Description: 	随机生成 token（生成的token已经自动在Redis里判重）
* @return string	生成的token字符串，其长度 = 20 + 4
* @return error
 */
func MakeToken(c *gin.Context) (string, error) {
	const Length = 20
	var token string
	for i := 0; i < 3; i++ {
		bytes := make([]byte, Length/2)
		if _, err := rand.Read(bytes); err != nil {
			log.Error("函数 MakeToken 生成随机字符串错误：" + err.Error())
			return "", err
		}
		token = "bbs-" + hex.EncodeToString(bytes)
		// 生成token后，下来是查询该 token 是否已经在token里（避免生成重复token）
		if isExist, err := isTokenExistInRedis(c, token); err != nil {
			log.Error("函数 isTokenExistInRedis 查询token是否存在报错：：" + err.Error())
			return "", err
		} else if isExist {
			// 该token已经存在，则重新生成
			continue
		} else {
			// 不存在，说明正常，则使用
			break
		}
	}
	if len(token) == 0 {
		msg := "函数 MakeToken 连续生成3次重复token："
		log.Error(msg)
		return "", errors.New(msg)
	}
	return token, nil
}

/*getToken
* @Description: 获取当前用户的token
* @param c
* @return string
* @return error
 */
func getToken(c *gin.Context) (string, error) {
	// 1、先拿到用户 token
	token := c.Request.Header.Get(config.Config.Common.HeaderTokenName)

	// 2. 如果 token 为空，返回错误
	if token == "" {
		return "", errors.New("登录过期，请登录")
	}
	return token, nil
}

/*isTokenExistInRedis
* @Description: 根据入参，去redis里查询该token是否存在
* @param token
 */
func isTokenExistInRedis(c *gin.Context, token string) (bool, error) {
	if res, err := db.RedisDB.Exists(c, token).Result(); err != nil {
		return false, err
	} else if res > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

/*SetBBSUserLoginByRedis
* @Description: 		将用户信息写入Redis。
*						写2个，第一个key是入参的token，value是用户信息。第二个key是根据用户ID生成的key，value是token
*						二者过期时间相同
* @param c
* @param token
* @param userData
* @return error
 */
func SetBBSUserLoginByRedis(c *gin.Context, token string, userData AdvanceBBSUserResponse) error {
	// 1. 需要两个键值对，一个是token作为key，另一个是用户ID作为key。他们互为键值对，方便互相反查
	userIDKey := getRedisKeyByUserID(userData.ID)

	// 2. 根据用户 userIDKey （因为这个是固定的），查询该值是否已经在 Redis 里
	if isExist, err := isTokenExistInRedis(c, userIDKey); err != nil {
		// 如果报错
		errMsg := "redis查询失败，key为：" + userIDKey
		log.Error(errMsg)
		return errors.New("redis查询失败")
	} else if isExist {
		// 如果已经存在，则根据值去反查另一个 Key，然后把另一个 Key 也删除掉
		anotherKey, errGetAnotherKey := db.RedisDB.Get(c, userIDKey).Result()
		if errGetAnotherKey != nil {
			errMsg := "redis查询失败，key为：" + userIDKey
			log.Error(errMsg)
			return errors.New("redis查询失败")
		}
		// 查询正常后，删除2个key，并继续后面的操作
		_, errDel := db.RedisDB.Del(c, userIDKey, anotherKey).Result()
		if errDel != nil {
			errMsg := "redis删除失败，key为：" + userIDKey + " 和 " + anotherKey
			log.Error(errMsg)
		}
	}

	// 3. 将这两个键值对写入 Redis
	// key的过期时间默认是7天，具体看 yml 配置
	var expireTime = config.Config.Common.LoginExpireTime * time.Hour
	if expireTime == 0 {
		expireTime = 24 * 7 * time.Hour
	}

	// 将序列化后的JSON数据存储到Redis中
	if err := db.RedisDB.Set(c, token, userIDKey, expireTime).Err(); err != nil {
		msg := fmt.Sprintf("用户信息写入redis失败，key为：%s，value为：%s", token, userIDKey)
		log.Error(msg)
		return errors.New(msg)
	}

	// 再以用户ID为 Key，将 token 作为值写入 redis，方便反查
	if err := db.RedisDB.Set(c, userIDKey, token, expireTime).Err(); err != nil {
		msg := fmt.Sprintf("用户信息写入redis失败，key为：%s，value为：%s", token, userIDKey)
		log.Error(msg)
		// 写入失败的话，则删除第一步操作
		db.RedisDB.Del(c, token)
		return errors.New(msg)
	}
	// 此时说明写入成功
	return nil
}

/*getRedisKeyByUserID
* @Description: 		根据规则和用户ID，生成对应的Redis的Key
* @param userID			用户ID
* @return string		生成的key
 */
func getRedisKeyByUserID(userID uint) string {
	// 再以用户ID为 Key，将 token 作为值写入 redis，方便反查
	userIDKey := fmt.Sprintf("AUTH-PC-BBS-USERID-%d", userID)
	return userIDKey
}

/*GetUserInfoFromToken
* @Description: 拿token，去redis里获取用户信息
* @param c
 */
func GetUserInfoFromToken(c *gin.Context) error {
	token, errGetToken := getToken(c)
	if errGetToken != nil {
		return errGetToken
	}

	// 3. 连接 redis，根据 redis key，获取对应的value
	redisKey := "bbs-" + token
	println(redisKey)

	return nil
}
