package userService

import (
	"main/db"
	"main/db/model"
)

/*GetUsersInfoByPage
* @Description: 返回对应页码的用户数据
* @param page		页码
* @return GetAllUsersInfoResponse	对应页码的数据
 */
func GetUsersInfoByPage(page uint) GetAllUsersInfoResponse {
	// 每页用户数目
	const UserPerPage = 100

	// 1. 先获取用户总数，根据单页数量，计算总共有多少页数据
	var totalUsers int64
	db.DbItem.Model(model.BBSUser{}).Count(&totalUsers)
	totalPages := (totalUsers + UserPerPage - 1) / UserPerPage

	// 如果请求页码的数据大于当前页码
	if page > uint(totalPages) {
		return GetAllUsersInfoResponse{
			Page:         page,
			TotalPage:    uint(totalPages),
			TotalUser:    uint(totalUsers),
			CountPerPage: UserPerPage,
			List:         make([]BaseBBSUserResponse, 0),
		}
	}

	// 2. 根据当前页码，返回对应页码的用户列表
	var users []model.BBSUser
	offset := (page - 1) * UserPerPage
	db.DbItem.Model(model.BBSUser{}).Offset(int(offset)).Limit(UserPerPage).Find(&users)

	var list []BaseBBSUserResponse
	for i := 0; i < len(users); i++ {
		var temp BaseBBSUserResponse
		temp.ConvertFromBBSUser(&users[i])
		list = append(list, temp)
	}

	// 3. 生成返回数据
	var res = GetAllUsersInfoResponse{
		Page:         page,
		TotalPage:    uint(totalPages),
		TotalUser:    uint(totalUsers),
		CountPerPage: UserPerPage,
		List:         list,
	}
	return res
}
