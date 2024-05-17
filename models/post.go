package models

import (
	"html/template"
	"time"
)

type Post struct {
	Pid        int           `json:"pid"`        // 文章ID
	Title      string        `json:"title"`      // 文章标题
	Slug       string        `json:"slug"`       // 文章别名 自定义页面 path
	Content    template.HTML `json:"content"`    // 文章内容 html
	Markdown   string        `json:"markdown"`   // 文章内容 markdown
	CategoryID int           `json:"categoryId"` // 文章分类ID
	UserID     int           `json:"userId"`     // 文章作者ID
	ViewCount  int           `json:"viewCount"`  // 文章浏览次数
	Type       int           `json:"type"`       // 文章类型 0:普通 1:自定义
	CreatedAt  time.Time     `json:"createdAt"`  // 文章创建时间
	UpdatedAt  time.Time     `json:"updatedAt"`  // 文章更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`          // 文章ID
	Title        string        `json:"title"`        // 文章标题
	Slug         string        `json:"slug"`         // 文章别名 自定义页面 path
	Content      template.HTML `json:"content"`      // 文章内容 html
	CategoryId   int           `json:"categoryId"`   // 文章分类ID
	CategoryName string        `json:"categoryName"` // 文章分类名称
	UserID       int           `json:"userId"`       // 文章作者ID
	UserName     string        `json:"userName"`     // 文章作者名称
	ViewCount    int           `json:"viewCount"`    // 文章浏览次数
	Type         int           `json:"type"`         // 文章类型 0:普通 1:自定义
	CreatedAt    string        `json:"createdAt"`    // 文章创建时间
	UpdatedAt    string        `json:"updatedAt"`    // 文章更新时间
}
