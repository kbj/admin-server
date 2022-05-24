// Package response 用户模块相关接口返回的实体
// @author: WeiKai
// @date: 2022/5/22
package response

// LoginResp 用户登录接口返回实体
type LoginResp struct {
	ID       *uint   `json:"id"`
	Name     *string `json:"name"`
	RealName *string `json:"realName"`
}
