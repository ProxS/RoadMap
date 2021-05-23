package main

import (
	"log"

	todo "github.com/ProxS/todo-app"
	"github.com/ProxS/todo-app/pkg/handler"
	"github.com/ProxS/todo-app/pkg/repository"
	"github.com/ProxS/todo-app/pkg/service"
	"github.com/spf13/viper"
	// "github.com/gin-gonic/gin"
)

// var port string = "8081"

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalf(`fatal error %s`, err)
		}
	}()

	if err := initConfig(); err != nil {
		log.Fatalf(`fatal error load config %s`, err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}

	// r := gin.New()
	// r.GET("/:message", func(c *gin.Context) {
	// 	message := c.Param("message")
	// 	c.JSON(200, gin.H{
	// 		"message": message,
	// 	})
	// })

	// r.Run(port) // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
