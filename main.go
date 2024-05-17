package main

import (
	"encoding/json"
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
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var indexData IndexData
	indexData.Title = "Hello World"
	indexData.Desc = "Hello World Desc"

	t := template.New("index.html")
	// 拿到html文件路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/resources/views/index.html")
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
