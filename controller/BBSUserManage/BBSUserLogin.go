package BBSUserManage

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"main/db/model"
	"main/handlers"
	"main/service/userService"
	"main/utils"
)

/*UserLogin
* @Description: 登录接口
* @param c
 */
func UserLogin(c *gin.Context) {
	var data userService.UserLoginRequest
	if err1 := c.ShouldBindJSON(&data); err1 != nil {
		// 这里必须调用一下翻译函数，错误情况下的返回 JSON 也是在这个方法里自动生成的
		log.Debug(err1)
		handlers.Translate(c, err1)
		return
	}

	// 1. 登录检查，错误则返回报错信息，成功则返回用户信息
	resUserData, errorCheck := isLoginSuccess(c, &data)
	if errorCheck != nil {
		utils.ErrorJson(c, errorCheck.Error())
		return
	}

	// 2. 根据用户信息，生成 token
	token, errMakeToken := userService.MakeToken(c)
	if errMakeToken != nil {
		// token生成错误
		log.Info("token生成错误，用户ID是：", resUserData.ID)
		utils.ErrorJson(c, "服务器错误")
		return
	}

	// 3. 将 Token 写入 redis 里
	if loginError := userService.SetBBSUserLoginByRedis(c, token, resUserData); loginError != nil {
		log.Info("token生成错误，用户ID是：", resUserData.ID)
		utils.ErrorJson(c, "登录失败，请联系管理员")
		return
	}

	// 4. 更新用户信息里的最后登录IP和登录地址
	userService.AfterBBSUserLoginSuccess(c, resUserData.ID)

	utils.SuccessJson(c, "登录成功", gin.H{
		"token":    token,
		"userInfo": resUserData,
	})
}

/*isLoginSuccess
* @Description: 		判断用户登录是否成功，主要是校验该用户账密、状态等，允许登录会返回true
* @param c
* @param loginData		账号和密码
* @return userService.AdvanceBBSUserResponse	登录成功会返回用户信息
* @return error			失败的话会返回错误信息
 */
func isLoginSuccess(c *gin.Context, loginData *userService.UserLoginRequest) (userService.AdvanceBBSUserResponse, error) {
	// 登录频率检查
	if isTooMany := userService.IsTooManyLogin(c); isTooMany {
		return userService.AdvanceBBSUserResponse{}, errors.New("登录频率过高，请稍后再尝试")
	}
	// 1. 先对密码进行加盐
	pwBySalt := userService.SaltAndHashPassword(loginData.Account, loginData.Password)

	// 2. 先账号调用查询用户的代码
	errGet, user := userService.GetBBSUserByAccount(loginData.Account)
	// 查询失败，一般是该用户不存在
	if errGet != nil {
		return userService.AdvanceBBSUserResponse{}, errGet
	}
	// 密码不同 说明密码错误
	if user.Password != pwBySalt {
		return userService.AdvanceBBSUserResponse{}, errors.New("密码错误")
	}
	// 如果 user.Status 不为 0，意味着用户状态错误，调用 GetStatusText(user.Status) 则返回错误提示信息
	if user.Status != model.UserStatusNormal {
		msg := fmt.Sprintf("用户状态错误：%s", user.GetBBSStatusText())
		return userService.AdvanceBBSUserResponse{}, errors.New(msg)
	}

	// 3. 将用户信息转为可以返回给用户的信息
	var resUserData userService.AdvanceBBSUserResponse
	resUserData.ConvertFromBBSUser(&user)
	return resUserData, nil
}
