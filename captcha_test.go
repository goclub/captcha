package xcaptcha_test

import (
	xcaptcha "github.com/goclub/captcha"
	"log"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	http.HandleFunc("/captcha", func(writer http.ResponseWriter, request *http.Request) {
		code, buffer, err :=  xcaptcha.New(xcaptcha.Option{}) ; if err != nil {
		    writer.WriteHeader(500)
		    log.Print(err)
		    return
		}
		http.SetCookie(writer, &http.Cookie{
			Name: "captcha_id",
			Value: "someid",
		})
		_, err = buffer.WriteTo(writer) ; if err != nil {
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
