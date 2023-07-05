package main

import (
	"github.com/aykutaslantas/gorm-crud/controllers"
	"github.com/aykutaslantas/gorm-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
initializers.LoadEnvVariables()
initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/post/:id", controllers.PostsGetOne)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("posts/:id", controllers.PostsDelete)
	r.Run()
}