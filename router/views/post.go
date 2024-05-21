package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/service"
)

func (*HTMLApi) PostDetail(w http.ResponseWriter, r *http.Request) {
	postView := common.Template.Detail
	path := r.URL.Path
	pIDStr := strings.TrimPrefix(path, "/p/")
	pIDStr = strings.TrimSuffix(pIDStr, ".html")
	pID, err := strconv.Atoi(pIDStr)
	if err != nil {
		postView.WriteError(w, errors.New("不识别此请求路径"))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		postView.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}

	res, err := service.GetPostsByID(pID)
	if err != nil {
		log.Println("文职详情获取数据错误:", err)
		postView.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	postView.WriteData(w, res)
}


