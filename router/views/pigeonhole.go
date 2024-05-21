package views

import (
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/service"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonholeView := common.Template.Pigeonhole

	pigeonholeRes, err := service.FindPosts()
	if err != nil {
		pigeonholeView.WriteError(w, err)
		return
	}
	pigeonholeView.WriteData(w, pigeonholeRes)
}
