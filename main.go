package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1]
}
func Date(layout string) string {
	return time.Now().Format(layout)
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
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(index, home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析模板出错...", err)
	}

	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "张三",
			ViewCount:    123,
			CreatedAt:    "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var homeRes = &responses.HomeRes{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	t.Execute(w, homeRes)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
