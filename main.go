package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var indexData IndexData
	indexData.Title = "Hello World"
	indexData.Desc = "Hello World Desc"

	t := template.New("index.html")
	// 拿到html文件路径
	path, _ := os.Getwd()
	index := path + "/resources/views/index.html"
	home := path + "/resources/views/home.html"
	header := path + "/resources/views/layout/header.html"
	footer := path + "/resources/views/layout/footer.html"
	personal := path + "/resources/views/layout/personal.html"
	post := path + "/resources/views/layout/post-list.html"
	pagination := path + "/resources/views/layout/pagination.html"
	t, _ = t.ParseFiles(index, home, header, footer, personal, post, pagination )
	t.Execute(w, indexData)
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
