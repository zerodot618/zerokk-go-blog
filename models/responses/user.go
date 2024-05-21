package responses

import "github.com/zerodot618/zerokk-go-blog/models"

type LoginRes struct {
	Token    string          `json:"token"`
	UserInfo models.UserInfo `json:"userInfo"`
}
