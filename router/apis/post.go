package apis

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/database"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/service"
	"github.com/zerodot618/zerokk-go-blog/utils"
)

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {
	// 获取用户
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
		return
	}
	uid := claims.Uid

	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId, _ := strconv.Atoi(params["categoryId"].(string))
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)

		post := &models.Post{
			Title:      title,
			Slug:       slug,
			Content:    template.HTML(content),
			Markdown:   markdown,
			CategoryID: cId,
			UserID:     uid,
			ViewCount:  0,
			Type:       int(postType),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err := service.SavePost(post)
		if err != nil {
			common.Error(w, err)
			return
		}
		common.Success(w, post)

	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		pId := params["pid"].(float64)
		cId, _ := params["categoryId"].(float64)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)

		post := &models.Post{
			Pid:        int(pId),
			Title:      title,
			Slug:       slug,
			Content:    template.HTML(content),
			Markdown:   markdown,
			CategoryID: int(cId),
			UserID:     uid,
			ViewCount:  0,
			Type:       int(postType),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		err := service.UpdatePost(post)
		if err != nil {
			common.Error(w, err)
			return
		}
		common.Success(w, post)
	}

}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIDStr := strings.TrimPrefix(path, "/api/v1/post/")
	pID, err := strconv.Atoi(pIDStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}

	post, err := database.GetPostsByID(pID)
	if err != nil {
		common.Error(w, err)
		return
	}

	common.Success(w, post)
}

func (*Api) SearchPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	searchRes, err := service.SearchPost(condition)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, searchRes)
}
