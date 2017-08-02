package main

import (
	"fmt"

	"github.com/KastuyaKawabe/go-quiz-api/router"
	"github.com/KatsuyaKawabe/go-quiz-api/db"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config error", err)
	}

}

func ListenAPI() {
	api := gin.New()
	api.Use(gin.Recovery())

	api.HandleMethodNotAllowed = true

	router.V2(api)
	router.HealthCheck(api)

	db.SetupDB()
	api.Run(":" + viper.GetString("api.port"))
}

func main() {
	ListenAPI()
}
