package views

import (
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/common"
	"github.com/zerodot618/zerokk-go-blog/service"
)

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writingView := common.Template.Writing
	wr, err := service.Writing()
	if err != nil {
		writingView.WriteError(w, err)
		return
	}
	writingView.WriteData(w, wr)
}
