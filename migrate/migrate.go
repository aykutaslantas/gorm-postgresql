package main

import (

	"github.com/aykutaslantas/gorm-crud/initializers"
	"github.com/aykutaslantas/gorm-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}