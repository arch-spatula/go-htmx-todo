package main

import (
	"log"
	"net/http"
	"strconv"

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
	var newId int

	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic("Db 연결에 실패하였습니다.")
	}

	db.AutoMigrate(&Todo{})

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		var todos []Todo

		db.Find(&todos)

		lastTodoIndex := len(todos) - 1
		if lastTodoIndex >= 0 {
			lastTodo := todos[lastTodoIndex]
			newId = lastTodo.Id
		}

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

		newId += 1

		ctx.HTML(http.StatusCreated, "todo-item.tmpl", gin.H{
			"Title":   todo.Title,
			"Content": todo.Context,
			"Id":      newId,
			"Done":    todo.Done,
		})
	})

	// 편집
	r.GET("/todo/edit/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return
		}

		var todo Todo
		result := db.First(&todo, idInt)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		ctx.HTML(http.StatusOK, "edit-item.tmpl", gin.H{
			"Title":   todo.Title,
			"Content": todo.Context,
			"Id":      newId,
			"Done":    todo.Done,
		})
	})

	// 갱신(저장)
	r.PATCH("/todo/:id", func(ctx *gin.Context) {
		//
		var todo Todo
		if ctx.ShouldBind(&todo) == nil {
			log.Print(todo.Id)
			log.Print(todo.Title)
			log.Print(todo.Context)
			log.Print(todo.Done)
		}

		// db.Update( , &Todo{Title: todo.Title, Context: todo.Context, Done: todo.Done, Id: todo.Id})

		ctx.HTML(http.StatusCreated, "todo-item.tmpl", gin.H{
			"Title":   todo.Title,
			"Content": todo.Context,
			"Id":      newId,
			"Done":    todo.Done,
		})
	})

	// 삭제
	r.DELETE("/todo/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
			return
		}

		var todo Todo
		result := db.First(&todo, idInt)
		if result.Error != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}

		result = db.Delete(&todo)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		ctx.HTML(http.StatusOK, "", nil)
	})

	r.Run()
}
