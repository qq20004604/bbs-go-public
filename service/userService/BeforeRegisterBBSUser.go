package userService

import "github.com/gin-gonic/gin"

/*BeforeRegisterBBSUser
* @Description: 注册前检查。例如检查该IP是否短时间内发起多次注册信息、是否注册了过多账号、该ip被ban掉等
*				正常的注册业务逻辑（例如检测账号/名字等是否重复，本函数不负责）
* @param c
* @return bool	true表示允许注册，false表示禁止注册
* @return error	系统级错误打log报错，这里抛出的error是自定义错误信息，用于返回给用户查看
 */
func BeforeRegisterBBSUser(c *gin.Context) (bool, error) {
	return true, nil
}
