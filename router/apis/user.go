package apis

import (
	"log"
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		log.Println("登录失败:", err)
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
