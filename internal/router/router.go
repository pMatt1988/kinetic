package router

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/pmatt1988/kinetic/internal/db"
	"github.com/pmatt1988/kinetic/internal/views"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.GET("/", func(ctx *gin.Context) {
		var people []db.Person
		db.Client.Find(&people)
		render(ctx, 200, views.Index(people))
	})

	router.PUT("/person", func(ctx *gin.Context) {
		name := ctx.PostForm("name")

		person := db.Person{Name: name}
		db.Client.Create(&person)
		render(ctx, 200, views.Person(person))
	})

	router.DELETE("/person/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		db.Client.Delete(&db.Person{}, id)

		ctx.Status(200)
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
