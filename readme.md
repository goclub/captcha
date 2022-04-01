# goclub/captcha

```shell
go get github.com/goclub/captcha
```

![demo](./demo.png)

```go
http.HandleFunc("/captcha", func(writer http.ResponseWriter, request *http.Request) {
    code, buffer, err :=  xcaptcha.New(xcaptcha.Option{}) ; if err != nil {
        writer.WriteHeader(500)
        log.Print(err)
        return
    }
    // 一定要先setcookie 再 WriteTo(writer)
    http.SetCookie(writer, &http.Cookie{
        Name: "captcha_id",
        Value: "someid",
    })
// 一定要先setcookie 再 WriteTo(writer)
    _, err = buffer.WriteTo(writer) ; if err != nil {
        writer.WriteHeader(500)
        log.Print(err)
        return
    }
    log.Print("code:", code)
    return
})
```
