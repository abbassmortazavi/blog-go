package routing

import (
	"blog/pkg/config"
	"log"
)

func Run() {
	router := GetRouter()
	configs := config.Get()
	err := router.Run(configs.Host + ":" + configs.Port)
	if err != nil {
		log.Fatal(err)
		return
	}
}
