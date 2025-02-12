package main

import (
	"blog/cmd"
	"blog/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	cmd.Execute()
}

func serve() {
	configs := configSet()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func configSet() config.Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading the configs")
	}

	var configs config.Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return configs
}
