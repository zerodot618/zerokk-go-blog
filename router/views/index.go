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

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败")
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10

	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	homeRes, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据错误:", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员！"))
	}
	index.WriteData(w, homeRes)
}
