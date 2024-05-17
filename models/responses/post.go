package responses

import (
	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models"
)

type SearchRes struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article models.PostMore
}
