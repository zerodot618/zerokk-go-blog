package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {

	t := template.New("index.html")
	// 拿到html文件路径
	path := config.Cfg.System.CurrentDir
	index := path + "/resources/views/index.html"
	home := path + "/resources/views/home.html"
	header := path + "/resources/views/layout/header.html"
	footer := path + "/resources/views/layout/footer.html"
	personal := path + "/resources/views/layout/personal.html"
	post := path + "/resources/views/layout/post-list.html"
	pagination := path + "/resources/views/layout/pagination.html"
	t, _ = t.ParseFiles(index, home, header, footer, personal, post, pagination)

	var homeRes = &responses.HomeRes{}
	t.Execute(w, homeRes)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
