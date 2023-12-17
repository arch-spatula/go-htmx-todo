package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ORM, 최근에 추가된 api 사용
// sqlite 데이터 베이스

type Todo struct {
	Id      int    `form:"id" gorm:"primaryKey"`
	Title   string `form:"title"`
	Context string `form:"content"`
	Done    bool   `form:"done"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		var todos []Todo

		results := db.Find(&todos)

		fmt.Println(results, todos)

		// 새로고침에 데이터 가져오기
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"heading": "hello htmx",
			"todos":   todos,
		})
	})

	r.POST("/todo", func(ctx *gin.Context) {

		var todo Todo
		if ctx.ShouldBind(&todo) == nil {
			log.Print(todo.Id)
			log.Print(todo.Title)
			log.Print(todo.Context)
			log.Print(todo.Done)
		}

		db.Create(&Todo{Title: todo.Title, Context: todo.Context, Done: todo.Done})

		ctx.HTML(http.StatusCreated, "todo-item.tmpl", gin.H{
			"title":   todo.Title,
			"content": todo.Context,
		})
	})

	// 편집

	// 갱신

	// 삭제

	r.Run()
}
