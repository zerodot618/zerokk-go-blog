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

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryView := common.Template.Category
	path := r.URL.Path
	cIDStr := strings.TrimPrefix(path, "/c/")
	cID, err := strconv.Atoi(cIDStr)
	if err != nil {
		categoryView.WriteError(w, errors.New("不识别此请求路径"))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		categoryView.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10

	res, err := service.GetPostsByCategoryID(page, pageSize, cID)
	if err != nil {
		log.Println("首页获取数据错误:", err)
		categoryView.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	categoryView.WriteData(w, res)
}
