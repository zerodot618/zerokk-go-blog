package service

import (
	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/database"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

func Writing() (wr responses.WritingRes, err error) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL

	categorys, err := database.GetAllCategory()
	if err != nil {
		return
	}
	wr.Categorys = categorys
	return
}
