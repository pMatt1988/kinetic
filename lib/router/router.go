package router

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/pmatt1988/kinetic/lib/db"
	"github.com/pmatt1988/kinetic/views"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		var people []db.Person
		db.Client.Find(&people)
		render(ctx, 200, views.Index(people))
	})

	router.POST("/add-person", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		db.Client.Create(&db.Person{Name: name})
		render(ctx, 200, views.Person(ctx.PostForm("name")))
	})

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}

	return router
}

func render(ctx *gin.Context, status int, template templ.Component) error {
	ctx.Status(status)
	return template.Render(ctx.Request.Context(), ctx.Writer)
}
