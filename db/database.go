package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/config"
	"main/db/model"
	"reflect"
)

var DbItem *gorm.DB
var RedisDB *redis.Client

func getDsn() string {
	dbConfig := config.Config.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)
	return dsn
}

func MySQLConnect() error {
	dsn := getDsn()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("MySQL连接失败， error=%v", err)
	}

	DbItem = db
	log.Info("MySQL连接成功：success")
	return nil
}

func hasDB(s interface{}) (bool, string) {
	v := reflect.ValueOf(s).Elem()
	tableName := v.MethodByName("TableName").Call([]reflect.Value{})[0].String()
	DbItem.AutoMigrate(&s)
	cResult := DbItem.Migrator().HasTable(tableName)
	return cResult, tableName
}

func RedisConnect() error {
	var Ctx = context.Background()
	options := redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: config.Config.Redis.Password,
		DB:       config.Config.Redis.DB, // use default DB
	}
	RedisDB = redis.NewClient(&options)
	if _, err := RedisDB.Ping(Ctx).Result(); err != nil {
		return err
	}
	log.Info("Redis 初始化成功")
	return nil
}

/*InitDatabase
* @Description:     初始化数据库，如果返回 false 则说明数据库初始化失败，直接结束服务
* @return error
**/
func InitDatabase() error {
	log.Info("初始化MySQL：开始")
	// 先进行数据库连接
	if connectError := MySQLConnect(); connectError != nil {
		return connectError
	}

	tableList := []interface{}{
		&model.BBSUser{},
		&model.Topic{},
	}

	log.Info("初始化MySQL表：开始执行")
	var errorMsg string
	for _, t := range tableList {
		if res, tableName := hasDB(t); res == false {
			errorMsg += fmt.Sprintf("%s 表并不存在或 %s 表创建失败\n", tableName, tableName)
		}
	}

	if len(errorMsg) > 0 {
		//log.Fatal("MySQL初始化：失败\n%s", errorMsg)
		return errors.New(errorMsg)
	}
	log.Info("初始化MySQL：成功")

	log.Info("初始化Redis：开始")
	if err := RedisConnect(); err != nil {
		return err
	}
	log.Info("初始化Redis：成功")

	return nil
}
