package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(&configurations)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

//this function is original
/*func configSet() config.Config {

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
}*/
