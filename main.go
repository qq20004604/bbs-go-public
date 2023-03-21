package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"main/config"
	"main/db"
	"main/middleware"
	"main/router"
	"main/scripts"
	"os"
	"strconv"
)

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
func GinHttps(isHttps bool) error {
	// ReleaseMode 模式，如果使用默认 debug 模式的话，注释掉这一行就行
	gin.SetMode(gin.ReleaseMode)
	//Test()
	dbError := db.InitDatabase()
	if dbError != nil {
		return dbError
	}
	r := gin.Default()

	r.GET("isLive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"live": true,
		})
	})

	//r.GET("/.well-known/pki-validation/fileauth.txt", func(c *gin.Context) {
	//	c.String(200, "20211021092032358xejrxcs74yzbvzlhaptx82qnop140n6uc3xxg1aizsihdic")
	//})

	if isHttps {
		r.Use(TlsHandler(7001))

		return r.RunTLS(":"+strconv.Itoa(7001), "./certificate/test.pem", "./certificate/test.key")
	}

	return r.Run(":" + strconv.Itoa(7001))
}
func GinHttp() error {
	// ReleaseMode 模式，如果使用默认 debug 模式的话，注释掉这一行就行
	gin.SetMode(gin.ReleaseMode)
	scripts.StartGinLog()
	r := gin.Default()
	r.Use(middleware.Recover)
	router.LoadRoute(r)

	log.Info("服务：运行中")
	if config.Config.Runtime.Env == "PROD" {
		if errRunProd := r.Run(":" + strconv.Itoa(7001)); errRunProd != nil {
			log.Fatal("服务：运行失败" + errRunProd.Error())
		}
	} else {
		if errRunLocal := r.Run("localhost:7001"); errRunLocal != nil {
			log.Fatal("服务：运行失败" + errRunLocal.Error())
		}
	}
	return nil
}

func Run() {
	if os.Getenv("ENVIRONMENT") == "PROD" {
		log.Info("当前环境：生产")
		os.Setenv("ENV", "PROD")
	} else {
		log.Info("当前环境：本地")
		log.Info("本地访问链接：http://127.0.0.1:7001")
		os.Setenv("ENV", "TEST")
	}
	scripts.BeforeRunServer()
	if GinHttp() != nil {
		return
	}
}

func main() {
	log.Info("服务：开始启动检测")
	Run()
	log.Info("服务：停止")
	//gitlabManageService.InsertGitProjectIntoGitProjectList()
}
