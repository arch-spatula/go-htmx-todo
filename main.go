package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 아래 코드 사용 금지
// http.NewServeMux()
// ORM, 최근에 추가된 api 사용금지
// sqlite 데이터 베이스

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello htmx",
		})
	})
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	r.Run()
}
