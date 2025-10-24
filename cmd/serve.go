package cmd

import (
	"blog/pkg/config"

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
	config.Set()
	configs := config.Get()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"app":     viper.GetString("APPNAME"),
		})
	})
	//both of them is correct
	//router.Run(viper.GetString("Host") + ":" + viper.GetString("Port"))
	router.Run(configs.Host + ":" + configs.Port)
}
