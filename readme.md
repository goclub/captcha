# goclub/captcha

```shell
go get github.com/goclub/captcha
```

![demo](./demo.png)

```go
http.HandleFunc("/captcha", func(writer http.ResponseWriter, request *http.Request) {
    code, err :=  xcaptcha.New(writer, xcaptcha.Option{}) 
    if err != nil {
        writer.WriteHeader(500)
        log.Print(err)
        return
    }
    log.Print("code:", code)
    return
})
```
