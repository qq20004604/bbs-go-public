package userService

import (
	"fmt"
	"main/db"
	"main/db/model"
)

/*BatchUpdateUserStatus
* @Description:
* @param c
* @param data
* @return error
 */
func BatchUpdateUserStatus(data BatchUpdateUserStatusRequest) error {
	// 拿到ID列表和状态
	list := data.List
	status := data.Status
	fmt.Println(list)

	// 批量更新 id 在 list 里面的数据，将其Status的值更新为 status
	err := db.DbItem.Model(&model.BBSUser{}).Where("id IN (?)", list).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}
