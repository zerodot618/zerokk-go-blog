package responses

import (
	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models"
)

type HomeRes struct {
	config.Viewer
	Categorys []models.Category
	Posts     []models.PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
