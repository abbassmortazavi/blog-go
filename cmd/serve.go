package cmd

import (
	"blog/config"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCMD)
}

var serveCMD = &cobra.Command{
	Use:   "serve",
	Short: "Serve App on development",
	Long:  "Application that serves the App on development",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	configSet()
	fmt.Println("Hello World")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"app":     viper.GetString("APPNAME"),
		})
	})
	router.Run(viper.GetString("Host") + ":" + viper.GetString("Port"))
}

func configSet() config.Config {

	//this setting for yaml file
	//viper.SetConfigName("config")
	//viper.SetConfigType("yml")

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var configs config.Config
	err = viper.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return configs
}
