package responses

import (
	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models"
)

type SearchRes struct {
	Pid   int    `json:"pid"` // 文章ID
	Title string `json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article models.PostMore
}

type WritingRes struct {
	Title     string
	CdnURL    string
	Categorys []models.Category
}

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []models.Category
	Lines     map[string][]models.Post
}
