package common

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/zerodot618/zerokk-go-blog/config"
	"github.com/zerodot618/zerokk-go-blog/models"
	"github.com/zerodot618/zerokk-go-blog/models/responses"
)

var Template models.HtmlTemplate

func LoadTemplate() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/resources/views/")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	return params
}

func Success(w http.ResponseWriter, data interface{}) {
	result := responses.Result{
		Code:  200,
		Error: "",
		Data:  data,
	}
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Error(w http.ResponseWriter, err error) {
	result := responses.Result{
		Code:  200,
		Error: err.Error(),
		Data:  "",
	}
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
