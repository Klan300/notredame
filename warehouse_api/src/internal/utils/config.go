package utils

import (
    "log"
    "strings"
    "fmt"

    "github.com/spf13/viper"
)

type ConfigDefinition struct {

    Source struct {
        Host string
        Database string
        Username string
        Password string
    }
    
    Target struct {
        Host string
    }

    Authen AuthenDefinition
}

type AuthenDefinition struct {
    Usernames []string
    Secret string
}

func (authen *AuthenDefinition) Exists (lookup string) bool {

    for _, username := range authen.Usernames {
        if username == lookup {
            return true
        }
    }
    return false
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

func Debug(message string, arguments ...interface{}){
    if strings.ToLower("DEBUG") == "debug" {
        if arguments != nil {
            message = fmt.Sprintf(message, arguments...)
        }
        log.Println(message)
    }
}

func Error(message string, arguments ...interface{}){
    if arguments != nil {
        message = fmt.Sprintf(message, arguments...)
    }
    log.Panicln(message)
}