package initRuntime

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

/*pathExists
* @Description: 判断目录是否存在
* @param path   路径
* @return bool  存在还是不存在
* @return error 报错
**/
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func dirCreate(path string) error {
	s, err := pathExists(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if s == false {
		os.Mkdir(path, os.ModePerm)
	}
	return nil
}

// InitDir 目录初始化
func InitDir() error {
	err1 := dirCreate("./log")

	if err1 != nil {
		return errors.New("目录初始化失败")
	} else {
		return nil
	}
}
