package server

import (
	"net/http"

	"github.com/zerodot618/zerokk-go-blog/router"
)

var App = &MsServer{}

type MsServer struct {
}

func (s *MsServer) Start(ip, port string) error {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
