package service

import (
	"errors"

	"github.com/zerodot618/zerokk-go-blog/database"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
	"github.com/zerodot618/zerokk-go-blog/utils"
)

func Login(userName, passwd string) (loginRes *responses.LoginRes, err error) {
	passwd = utils.Md5Crypt(passwd, "zerokk")
	user, err := database.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	// 生成 token
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token 未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	userInfo.CreatedAt = user.CreatedAt
	userInfo.UpdatedAt = user.UpdatedAt
	loginRes = &responses.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}
	return loginRes, nil
}
