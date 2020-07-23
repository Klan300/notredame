package utils

import (
    "log"
    "time"

    "github.com/spf13/viper"
)

type ConfigDefinition struct {
    Source struct {
        Host string
        Token string
        Consumers int
        Attempts int
        Wait time.Duration
    }

    Target struct {
        Host string
        Database string
        Username string
        Password string
    }

    Exchanges []string
    Documents []string

    Logging struct {
        Level string
        Stdout bool
        Dirname string
    }
}


var Config *ConfigDefinition = new(ConfigDefinition)
func LoadConfig(filePath string){

    viper.AutomaticEnv()
    viper.SetConfigFile(filePath)

    err := viper.ReadInConfig()
    if err != nil {
        log.Panicf("[config.go] %v\n", err)
    }

    viper.Unmarshal(&Config)
    LoadLoggingConfig()
}

