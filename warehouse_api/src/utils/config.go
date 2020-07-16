package utils

import (
    "log"

    "github.com/spf13/viper"
)

type ConfigDefinition struct {
	
	Usernames []string

	Secret string

	Source struct {
		Host string
        Database string
        Username string
        Password string
	} 
}


var Config *ConfigDefinition

func LoadConfig(filePath string){

    viper.AutomaticEnv()
    viper.SetConfigFile(filePath)

    err := viper.ReadInConfig()
    if err != nil {
        log.Panicf("[config.go] %s\n", err)
    }

	viper.Unmarshal(&Config)

}