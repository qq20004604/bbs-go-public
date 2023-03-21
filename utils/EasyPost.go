package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// EasyPostWithToken
// @Description: 发起一个带 token 的 post 请求
// @param c
// @param inface    数据
// @return interface{}  该结构体
// @return error
func EasyPostWithToken(c *gin.Context, url string, inface interface{}) error {
	// 1、先拿到用户 token
	token := c.Request.Header.Get("token")

	// 2. 如果 token 为空，返回错误
	if token == "" {
		return errors.New("请登录后再进行操作")
	}

	// 2、再发起异步请求，获取当前用户信息
	data := make(map[string]interface{})
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST",
		url,
		bytes.NewReader(bytesData))
	req.Header.Set("token", token)
	response, err := (&http.Client{}).Do(req)
	defer response.Body.Close()

	// 如果请求报错
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		msg := "服务调用错误"
		log.Error(url + "接口调用失败")
		return errors.New(msg)
	}
	body, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &inface)

	return nil
}

// EasyPostWithTokenAndData
// @Description: 发起一个带 token ，并且带参数的的 post 请求
// @param c
// @param inface    数据
// @return interface{}  该结构体
// @return error
func EasyPostWithTokenAndData(c *gin.Context, url string, inface interface{}, data map[string]interface{}) error {
	// 1、先拿到用户 token
	token := c.Request.Header.Get("token")
	// token 为空则返回
	if token == "" {
		msg := "请登录后再进行操作"
		//ErrorJson(c, msg)
		return errors.New(msg)
	}

	// 2、再发起异步请求，获取当前用户信息
	bytesData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST",
		url,
		bytes.NewReader(bytesData))
	req.Header.Set("token", token)
	response, err := (&http.Client{}).Do(req)
	defer response.Body.Close()

	// 如果请求报错
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		msg := "服务调用错误"
		log.Error(url + "接口调用失败")
		return errors.New(msg)
	}
	body, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(body, &inface)

	return nil
}
