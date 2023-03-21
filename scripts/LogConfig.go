package scripts

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"io"
	"main/config"
	"main/utils"
	"os"
	"path"
	"time"
)

func ConfigLocalFilesystemLogger(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	baseLogPath := path.Join(logPath, logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+"-%Y%m%d%H%M.log",
		//rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间（超出这个时间的日志文件会被删除）
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local fileManage system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer, // 为不同级别设置不同的输出目的
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})
	// 将函数名和行数放在日志里面
	log.SetReportCaller(config.Config.Runtime.SetReportCaller)
	log.AddHook(lfHook)
}
func InitLogfile() {
	s, err := os.Stat("./log")
	if err != nil {
		return
	}
	if s.IsDir() == false {
		// 判断该文件不是文件夹，于是创建文件夹
		err2 := os.Mkdir("./log", os.ModePerm)
		if err2 != nil {
			fmt.Println(err2)
		}
		log.Info("log 文件夹：创建完毕")
	} else {
		log.Info("log 文件夹：已存在")
	}
}

/*InitLog
* @Description: 初始化 Log
**/
func InitLog() {
	InitLogfile()
	log.SetLevel(config.Config.Runtime.LogLevel)
	exPath := utils.GetCurrentAbPath()
	// 进行配置日志目录
	p := path.Join(exPath, "/../log")
	//log.Info(p)
	ConfigLocalFilesystemLogger(p, "logfile", time.Second*60*60*24*config.Config.Runtime.LogMaxAge, time.Second*60*60*config.Config.Runtime.LogRotationTime)
}

func StartGinLog() {
	// Gin 本身的日志启用
	//log.Info("——————————", utils.GetCurrentAbPath())
	_, err1 := os.Stat("./log/gin.log")
	if err1 != nil {
		os.Create("log/gin.log")
	}
	f, err3 := os.OpenFile("log/gin.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err3 != nil {
		log.Error(err3)
		return
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
