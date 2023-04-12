package main

import (
	"github.com/Geo-Bit/go-api/initializers"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
