package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"main/utils"
	"os"
	"time"
)

// YMLStruct
// @Description: YML 的数据结构
type YMLStruct struct {
	MySQL struct {
		Default MsSQLConfig `yaml:"default"`
		DEV     MsSQLConfig `yaml:"dev"`
		PROD    MsSQLConfig `yaml:"prod"`
	} `yaml:"mysql"`
	Redis struct {
		Default RedisConfig `yaml:"default"`
		DEV     RedisConfig `yaml:"dev"`
		PROD    RedisConfig `yaml:"prod"`
	} `yaml:"redis"`
	Common  CommonConfig `yaml:"commonConfig"`
	Runtime struct {
		Default RuntimeConfig `yaml:"default"`
		DEV     RuntimeConfig `yaml:"dev"`
		PROD    RuntimeConfig `yaml:"prod"`
	} `yaml:"runtimeConfig"`
}

// Configuration
// @Description: 读取后存储的数据结构
type Configuration struct {
	MySQL   MsSQLConfig
	Redis   RedisConfig
	Common  CommonConfig
	Runtime RuntimeConfig
}

// MsSQLConfig
// @Description: MySQL 的配置
type MsSQLConfig struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Host     string `yaml:"Host,omitempty"`
	Port     int    `yaml:"Port,omitempty"`
	Dbname   string `yaml:"Dbname,omitempty"`
}

// RedisConfig
// @Description: Redis的配置
type RedisConfig struct {
	Addr     string `yaml:"addr,omitempty"`
	Password string `yaml:"password,omitempty"`
	DB       int    `yaml:"db,omitempty"`
}

/*RuntimeConfig
* @Description: 运行时设置，主要是一些系统配置，例如日志是否显示，当前是什么环境之类，通常只会获取一次。
*				 具体说明参考yml文件
 */
type RuntimeConfig struct {
	Env                string        `yaml:"Env,omitempty"`
	SetReportCaller    bool          `yaml:"SetReportCaller,omitempty"`
	ServerName         string        `yaml:"ServerName,omitempty"`
	BaseUrl            string        `yaml:"BaseUrl,omitempty"`
	ServerURL          string        `yaml:"ServerURL,omitempty"`
	LogMaxAge          time.Duration `yaml:"LogMaxAge,omitempty"`
	LogRotationTime    time.Duration `yaml:"LogRotationTime,omitempty"`
	LogLevel           log.Level     `yaml:"LogLevel"`
	CreateFirstAdmin   bool          `yaml:"CreateFirstAdmin,omitempty"`
	FirstAdminAccount  string        `yaml:"FirstAdminAccount,omitempty"`
	FirstAdminName     string        `yaml:"FirstAdminName,omitempty"`
	FirstAdminPassword string        `yaml:"FirstAdminPassword,omitempty"`
}

// CommonConfig
// @Description: 通用设置
type CommonConfig struct {
	PasswordSalt            string `yaml:"PasswordSalt"`
	PasswordLengthAfterHash uint   `yaml:"PasswordLengthAfterHash"`
	HeaderTokenName         string `yaml:"HeaderTokenName"`
}

// Config 全局变量
var Config Configuration

// loadConfiguration
//
//	@Description: YML加载配置
//	@param filePath
//	@return error
func loadConfiguration() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("YML加载失败：捕获错误如下，", r)
		}
	}()

	file, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal("YML加载失败：文件打开失败。", err.Error())
	}

	var ymlData YMLStruct
	if err = yaml.Unmarshal(file, &ymlData); err != nil {
		log.Fatal("YML加载失败：序列化解析失败。", err.Error())
	}

	if os.Getenv("ENV") == "PROD" {
		utils.MergeSettings(&Config.MySQL, &ymlData.MySQL, "Default", "PROD")
		utils.MergeSettings(&Config.Redis, &ymlData.Redis, "Default", "PROD")
		utils.MergeSettings(&Config.Runtime, &ymlData.Runtime, "Default", "PROD")
		utils.MergeData(&Config.Common, &ymlData.Common)
	} else {
		utils.MergeSettings(&Config.MySQL, &ymlData.MySQL, "Default", "DEV")
		utils.MergeSettings(&Config.Redis, &ymlData.Redis, "Default", "DEV")
		utils.MergeSettings(&Config.Runtime, &ymlData.Runtime, "Default", "DEV")
		utils.MergeData(&Config.Common, &ymlData.Common)
	}
}

// LoadYML
// @Description: 读取YML配置
// @return error
func LoadYML() {
	log.Info("YML配置：开始加载")
	loadConfiguration()
	log.Info("YML配置：加载成功")
}
