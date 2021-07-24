package xcaptcha_test

import (
	xcaptcha "github.com/goclub/captcha"
	"log"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	http.HandleFunc("/captcha", func(writer http.ResponseWriter, request *http.Request) {
		code, err :=  xcaptcha.New(writer, xcaptcha.Option{}) ; if err != nil {
		    writer.WriteHeader(500)
		    log.Print(err)
		    return
		}
		log.Print("code:", code)
		return
	})
	addr := ":9241"
	log.Print("http://127.0.0.1" + addr + "/captcha")
	log.Print(http.ListenAndServe(addr, nil))
}
