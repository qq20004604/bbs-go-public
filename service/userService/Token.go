package userService

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db"
	"regexp"
	"strconv"
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
		if err == redis.Nil {
			// 说明键值对不存在
			return false, nil
		} else {
			log.Error("redis查询错误")
		}
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
	if errOld != nil {
		// errOld == redis.Nil 为true时，这里说明键值对不存在，即 userIDKey 这个key不存在redis里，所以不需要删除
		// 而为 errOld != redis.Nil时，说明是其他错误（例如连接错误、超时等）
		if errOld != redis.Nil {
			// 此时，打印错误日志
			log.Error(errOld)
		}
	} else {
		// 当这个报错不存在时，说明键值对存在，所以直接都删
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
		// 这个说明是其他错误
		if errGetAnotherKey != redis.Nil {
			log.Error("redis查询错误")
		}
	} else {
		// 查询正常后，删除2个key，并继续后面的操作
		db.RedisDB.Del(c, token, anotherKey)
	}
}

/*CheckTokenAvailable
* @Description: 检查 token 是否有效
* @param c
* @param token
* @return error
 */
func CheckTokenAvailable(c *gin.Context, token string) (error, uint) {
	// 从 Redis 中查询 token
	userIDKey, err := db.RedisDB.Get(c, token).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("未登录或登录过期，请重新登录"), 0
		}
		log.Error(err)
		return errors.New("服务器错误"), 0
	}

	// 以查询到的值作为 key 再次查询 Redis
	redisToken, err := db.RedisDB.Get(c, userIDKey).Result()
	if err != nil {
		if err == redis.Nil {
			db.RedisDB.Del(c, token)
			return errors.New("未登录或登录过期，请重新登录"), 0
		}
		log.Error(err)
		return errors.New("服务器错误"), 0
	}

	// 比对查询到的值和 token 是否一致
	if redisToken != token {
		db.RedisDB.Del(c, token)
		return errors.New("未登录或登录过期，请重新登录"), 0
	}
	re := regexp.MustCompile(`AUTH-PC-BBS-USERID-(\d+)`)
	match := re.FindStringSubmatch(userIDKey)
	if len(match) > 1 {
		idStr := match[1]
		id, errMatch := strconv.ParseUint(idStr, 10, 64)
		if errMatch != nil {
			// 匹配失败，一般是数字太大，超过 uint 限制
			log.Error(fmt.Println("Error parsing number:", err))
			return errors.New("服务器错误"), 0
		} else {
			// 理论上，数字太大然后转失败了，需要额外考虑这种情况，但实际上，应该问题不大，毕竟 uint 已经很大了，而且也写不进 redis 里
			return nil, uint(id)
		}
	} else {
		// 理论上，不太可能出现这种情况，即 token 和 userIDKey 互相匹配存在 redis 里，但匹配不到数据
		// 只有一种可能，生成 userIDKey 和 获取的正则 不匹配。
		log.Error("请检查生成 userIDKey 的规则 和 正则匹配的表达式，其是否一致")
		return errors.New("服务器错误，请联系管理员"), 0
	}
}
