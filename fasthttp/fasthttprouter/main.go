package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func httpHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello,%v",ctx.UserValue("name"))
}

func httpHandler1(ctx *fasthttp.RequestCtx)  {
	ctx.SetContentType("text/html")
	fmt.Fprintf(ctx, "Method:%s <br/>", ctx.Method())
	fmt.Fprintf(ctx, "URI:%s <br/>", ctx.URI())
	fmt.Fprintf(ctx, "RemoteAddr:%s <br/>", ctx.RemoteAddr())
	fmt.Fprintf(ctx, "UserAgent:%s <br/>", ctx.UserAgent())
	fmt.Fprintf(ctx, "Header.Accept:%s <br/>", ctx.Request.Header.Peek("Accept"))
}

func httpHandler2(ctx *fasthttp.RequestCtx)  {
	ctx.SetContentType("text/html")
	fmt.Fprintf(ctx, "IP:%s <br/>", ctx.RemoteIP())
	fmt.Fprintf(ctx, "Host:%s <br/>", ctx.Host())
	fmt.Fprintf(ctx, "ConnectTime:%s <br/>", ctx.ConnTime()) // 连接收到处理的时间
	fmt.Fprintf(ctx, "IsGET:%v <br/>", ctx.IsGet())          // 类似有 IsPOST, IsPUT 等
}

func main() {
	router := fasthttprouter.New()
	router.GET("/hello/:name", httpHandler)
	router.GET("/hello1",httpHandler1)
	router.GET("/hello2",httpHandler2)
	if err := fasthttp.ListenAndServe(":8089", router.Handler); err != nil {
		fmt.Println("start fasthttp fail:", err)
	}
}
