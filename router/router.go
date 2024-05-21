package router

import (
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/router/apis"
	"github.com/zerodot618/zerokk-go-blog/router/views"
)

func Router() {
	// 页面
	http.HandleFunc("/", views.HTML.Index)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/p/", views.HTML.PostDetail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)

	// api 数据
	http.HandleFunc("/api/v1/post", apis.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", apis.API.GetPost)
	http.HandleFunc("/api/v1/post/search", apis.API.SearchPost)
	http.HandleFunc("/api/v1/qiniu/token", apis.API.QiniuToken)
	http.HandleFunc("/api/v1/login", apis.API.Login)

	// 静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
