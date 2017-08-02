package main

import (
	"fmt"

	"github.com/KatsuyaKawabe/go-quiz-api/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("dev")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config error", err)
	}

}

func ListenAPI() {
	api := gin.New()
	api.Use(gin.Recovery())

	api.HandleMethodNotAllowed = true

	router.V1(api)
	api.Run(":" + viper.GetString("api.port"))
}

func main() {
	ListenAPI()
}
