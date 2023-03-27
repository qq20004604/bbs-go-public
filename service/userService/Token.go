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
	oldToken, errOld := db.RedisDB.Get(c, userIDKey).Result()

	// 3. 将这两个键值对写入 Redis
	// 创建一个新的 redis 事务，合并操作，减少bug出现的概率
	pipe := db.RedisDB.TxPipeline()
	// 查不到的话则不删除
	if errOld == nil {
		// 删除之前的两个
		pipe.Del(c, oldToken, userIDKey)
	}
	// 使用 MSet 方法一次性设置两个键值对
	pipe.MSet(c, token, userIDKey, userIDKey, token)
	// key的过期时间默认是7天，具体看 yml 配置
	var expireTime = config.Config.Common.LoginExpireTime * time.Hour
	if expireTime == 0 {
		expireTime = 24 * 7 * time.Hour
	}
	// 为第一个键（token）设置过期时间
	pipe.Expire(c, token, expireTime)
	// 为第二个键（userIDKey）设置过期时间
	pipe.Expire(c, userIDKey, expireTime)
	// 执行事务
	if _, err := pipe.Exec(c); err != nil {
		msg := fmt.Sprintf("用户信息写入redis失败，key为：%s，value为：%s", token, userIDKey)
		log.Error(msg)
		log.Error(err)
		db.RedisDB.Del(c, token, userIDKey)
		return errors.New(msg)
	}

	// 4. 将 token 再写入 cookie 里，默认过期时间（14天）
	SetLoginByCookie(c, token)

	// 6. 说明正常，返回用户信息
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

/*ClearTokenByRedis
* @Description:		清除 redis 里的 token 信息，同时清除 token 和 用户ID 两组key
* @param token		入参是 token，或者用户ID的key，都可以
 */
func ClearTokenByRedis(c *gin.Context, token string) {
	// 如果已经存在，则根据值去反查另一个 Key，然后把另一个 Key 也删除掉
	anotherKey, errGetAnotherKey := db.RedisDB.Get(c, token).Result()
	if errGetAnotherKey != nil {
		db.RedisDB.Del(c, token)
	} else {
		// 查询正常后，删除2个key，并继续后面的操作
		db.RedisDB.Del(c, token, anotherKey)
	}
}
