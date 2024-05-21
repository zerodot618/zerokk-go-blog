package views

import (
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/config"
)

func (*HTMLApi) Login(w http.ResponseWriter, r *http.Request) {
	loginView := common.Template.Login

	loginView.WriteData(w, config.Cfg.Viewer)
}
