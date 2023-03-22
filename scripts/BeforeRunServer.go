package scripts

import (
	log "github.com/sirupsen/logrus"
	"main/config"
	"main/db"
	"main/handlers"
	"main/initRuntime"
)

func BeforeRunServer() {
	// 加载配置文件（如果报错会自动退出程序）
	config.LoadYML()
	//e := GinHttps(true) // 这里false 表示 http 服务，非 https
	if initRuntime.InitDir() != nil {
		return
	}
	// 初始化日志设置
	InitLog()
	// 注册校验组件中文翻译器
	handlers.RegisterCNValidator()
	// 初始化数据库
	dbError := db.InitDatabase()
	if dbError != nil {
		log.Fatal(dbError.Error())
		return
	}
	// 创建超级管理员账号
	CreateFirstSuperAdmin()
}
